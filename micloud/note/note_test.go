package notemgr

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"testing"
)

func TestListNotes(t *testing.T) {
	n := Note{
		Client:        client.GlobalClient,
		NumOfReqInSec: 5,
	}
	notes, err := n.ListNotes(200)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(notes)
}

func TestGetNote(t *testing.T) {
	n := Note{
		Client:        client.GlobalClient,
		NumOfReqInSec: 5,
	}
	note, err := n.GetNote("43465194589668640")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(note)
}

func TestDeleteNote(t *testing.T) {
	n := Note{
		Client:        client.GlobalClient,
		NumOfReqInSec: 5,
	}
	err := n.DeleteNote("43465194589668640", "43465212726624768", false)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateNote(t *testing.T) {
	n := Note{
		Client:        client.GlobalClient,
		NumOfReqInSec: 5,
	}
	notes, err := n.ListDeletedNotes("43465212726624768", 200)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(notes)
}
