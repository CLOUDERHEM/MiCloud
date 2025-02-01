package contact

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"strconv"
	"time"
)

const (
	initData = "https://i.mi.com/contacts/initdata"
)

func ListContacts(client *client.Client, limit int) (Contacts, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"_dc", ts},
		{"syncTag", "0"},
		{"limit", strconv.Itoa(limit)},
		{"syncIgnoreTag", "0"},
	}

	body, r, err := client.DoRequest(request.NewGet(initData, q))
	if err != nil {
		return Contacts{}, err
	}
	return validate.Validate[Contacts](r, body)
}
