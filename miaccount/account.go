package miaccount

import (
	"errors"
	"github.com/clouderhem/micloud/utility/parse"
	"github.com/clouderhem/micloud/utility/request"
	"strings"
	"time"
)

func New(miAccountCookie string) *MiAccount {
	return &MiAccount{
		Cookie:  miAccountCookie,
		Timeout: 5 * time.Second,
	}
}

type MiAccount struct {
	Cookie  string
	Timeout time.Duration
}

func (m *MiAccount) GenMiCloudCookie() (string, error) {
	serviceLoginUrl, err := GetLoginUrl()
	if err != nil {
		return "", err
	}
	stsUrl, err := m.getSTSUrl(serviceLoginUrl)
	if err != nil {
		return "", err
	}
	return m.getMiCloudCookie(stsUrl)
}

func (m *MiAccount) GetServiceToken() string {
	return parse.GetValueByKey(m.Cookie, "serviceToken")
}

func (m *MiAccount) getSTSUrl(serviceLoginUrl string) (string, error) {
	req := request.NewGet(serviceLoginUrl, nil)

	req.Header.Add("Cookie", m.Cookie)
	_, resp, err := request.DoRequest(req, request.Timeout(m.Timeout))
	if err != nil {
		return "", err
	}
	location := resp.Header.Get("Location")
	if location == "" {
		return "", errors.New("no location in service login resp")
	}
	return location, nil
}

func (m *MiAccount) getMiCloudCookie(stsUrl string) (string, error) {
	req := request.NewGet(stsUrl, nil)
	req.Header.Set("Cookie", m.Cookie)
	_, resp, err := request.DoRequest(req, request.Timeout(m.Timeout))
	if err != nil {
		return "", err
	}

	values := resp.Header.Values("Set-Cookie")
	if len(values) == 0 {
		return "", errors.New("no cookies in sts resp")
	}
	return parse.TidyKvs(strings.Join(values, ";")), nil
}
