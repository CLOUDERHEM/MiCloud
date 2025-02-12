package request

import (
	"io"
	"net/http"
)

type UrlQuery struct {
	Key   string
	Value string
}

func DoRequest(req *http.Request, opts ...Opt) (body []byte, resp *http.Response, err error) {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return doRequest(req, c)
}

func NewGet(url string, queries []UrlQuery) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query()
	for _, kv := range queries {
		q.Add(kv.Key, kv.Value)
	}
	req.URL.RawQuery = q.Encode()

	return req
}

func doRequest(req *http.Request, c *Config) (body []byte, resp *http.Response, err error) {
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	client.Timeout = c.Timeout
	defer client.CloseIdleConnections()

	resp, err = client.Do(req)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, resp, err
}
