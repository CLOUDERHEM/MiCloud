package contactmgr

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/contact/contact"
)

type Contact struct {
	Client *client.Client
}

func (c *Contact) ListContacts(limit int) (contact.Contacts, error) {
	return contact.ListContacts(c.Client, limit)
}
