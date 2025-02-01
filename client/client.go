package client

import (
	"github.com/clouderhem/micloud/miaccount"
	"github.com/clouderhem/micloud/utility/request"
	"log"
	"net/http"
	"time"
)

var GlobalCookie = ""

var GlobalClient = &Client{
	MiAccount: &miaccount.MiAccount{
		Cookie:  GlobalCookie,
		Timeout: 5 * time.Second,
	},
	RetryTimes: 1,
	Timeout:    5 * time.Second,
}

func New(mi *miaccount.MiAccount) *Client {
	return &Client{
		MiAccount:  mi,
		RetryTimes: 2,
		Timeout:    time.Second * 5,
	}
}

type Client struct {
	MiAccount *miaccount.MiAccount

	RetryTimes int
	Timeout    time.Duration

	cookie string
}

func (c *Client) DoRequest(req *http.Request) ([]byte, *http.Response, error) {
	var body []byte
	var resp *http.Response
	var err error
	for i := 0; i < c.RetryTimes+1; i++ {
		c.postProcessReq(req)
		body, resp, err = request.DoRequest(req, request.Timeout(c.Timeout))
		if err != nil {
			continue
		}
		if resp.StatusCode == http.StatusUnauthorized {
			cookie, err := c.MiAccount.GenMiCloudCookie()
			if err != nil || cookie == "" {
				log.Print("cannot renewal micloud cookie, err: ", err)
				return body, resp, err
			} else {
				c.cookie = cookie
			}
			continue
		}
		break
	}

	return body, resp, err
}

func (c *Client) postProcessReq(req *http.Request) {
	req.Header.Set("Cookie", c.cookie)
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
}
