package micloud

import (
	"github.com/clouderhem/micloud/miaccount"
	"testing"
)

func TestNew(t *testing.T) {
	cloud := New(miaccount.New("your_mi_account_cookie"))

	devices, err := cloud.Device.ListDevices()
	if err != nil {
		t.Error(err)
	}
	for _, device := range devices {
		t.Log(device)
	}
}
