package devicemgr

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"testing"
)

func TestListDevices(t *testing.T) {
	d := Device{
		Client: client.GlobalClient,
	}
	devices, err := d.ListDevices()
	if err != nil {
		t.Error(err)
	}
	if len(devices) == 0 {
		t.Error("no devices found")
	}
}

func TestGetDeviceStatus(t *testing.T) {
	d := Device{
		Client: client.GlobalClient,
	}
	status, err := d.GetDeviceStatus()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(status)
}
