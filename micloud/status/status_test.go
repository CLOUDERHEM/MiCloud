package statusmgr

import (
	"github.com/clouderhem/micloud/client"
	"testing"
)

func TestGetDetail(t *testing.T) {
	s := Status{
		Client: client.GlobalClient,
	}
	detail, err := s.GetDetail()
	if err != nil {
		t.Error(err)
	}
	if detail.SettingType == "" {
		t.Error("setting type is empty")
	}
}

func TestRenewal(t *testing.T) {
	s := Status{
		Client: client.GlobalClient,
	}
	renewal, err := s.Renewal()
	if err != nil {
		t.Error(err)
	}
	if renewal == "" {
		t.Error("no cookie found")
	}
}
