package smsmgr

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/sms/message"
	"github.com/clouderhem/micloud/micloud/sms/recyclebin"
)

type Sms struct {
	Client *client.Client
}

func (s *Sms) ListMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	return message.ListMessages(s.Client, syncTag, syncThreadTag, limit)
}

func (s *Sms) ListDeletedMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	return recyclebin.ListDeletedMessages(s.Client, syncTag, syncThreadTag, limit)
}
