package statusmgr

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/status/detail"
	"github.com/clouderhem/micloud/micloud/status/setting"
)

type Status struct {
	Client *client.Client
}

// GetDetail get micloud service detail
func (s *Status) GetDetail() (detail.Detail, error) {
	return detail.GetAllDetail(s.Client)
}

// Renewal renew cookie
func (s *Status) Renewal() (string, error) {
	cookie, err := setting.Renewal(s.Client)
	if err != nil {
		return "", err
	}
	return cookie, err
}
