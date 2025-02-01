package contactmgr

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"testing"
)

func TestListContacts(t *testing.T) {
	c := Contact{
		Client: client.GlobalClient,
	}
	contacts, err := c.ListContacts(200)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(contacts)
}
