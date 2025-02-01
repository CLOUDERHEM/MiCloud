package status

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"time"
)

const (
	deviceStatusApi = "https://i.mi.com/find/device/full/status?ts=%v"
)

func GetDeviceStatus(client *client.Client) (Status, error) {
	req := request.NewGet(fmt.Sprintf(deviceStatusApi, time.Now().UnixMilli()), nil)
	body, r, err := client.DoRequest(req)
	if err != nil {
		return Status{}, err
	}

	return validate.Validate[Status](r, body)
}
