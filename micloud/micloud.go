package micloud

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/miaccount"
	contactmgr "github.com/clouderhem/micloud/micloud/contact"
	devicemgr "github.com/clouderhem/micloud/micloud/device"
	gallerymgr "github.com/clouderhem/micloud/micloud/gallery"
	notemgr "github.com/clouderhem/micloud/micloud/note"
	recordingmgr "github.com/clouderhem/micloud/micloud/recording"
	smsmgr "github.com/clouderhem/micloud/micloud/sms"
	statusmgr "github.com/clouderhem/micloud/micloud/status"
)

func New(m *miaccount.MiAccount) *MiCloud {
	c := client.New(m)
	return &MiCloud{
		Contact:   &contactmgr.Contact{Client: c},
		Device:    &devicemgr.Device{Client: c},
		Gallery:   &gallerymgr.Gallery{Client: c},
		Note:      &notemgr.Note{Client: c, NumOfReqInSec: 5},
		Recording: &recordingmgr.Recording{Client: c},
		Sms:       &smsmgr.Sms{Client: c},
		Status:    &statusmgr.Status{Client: c},
	}
}

type MiCloud struct {
	Contact   *contactmgr.Contact
	Device    *devicemgr.Device
	Gallery   *gallerymgr.Gallery
	Note      *notemgr.Note
	Recording *recordingmgr.Recording
	Sms       *smsmgr.Sms
	Status    *statusmgr.Status
}
