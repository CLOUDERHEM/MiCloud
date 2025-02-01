package devicemgr

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/device/device"
	"github.com/clouderhem/micloud/micloud/device/status"
)

type Device struct {
	Client *client.Client
}

// ListDevices list all devices logged in with current xiaomi account
func (d *Device) ListDevices() ([]device.Device, error) {
	return device.ListDevices(d.Client)
}

// GetDeviceStatus list all devices' status including position info
func (d *Device) GetDeviceStatus() (status.Status, error) {
	return status.GetDeviceStatus(d.Client)
}
