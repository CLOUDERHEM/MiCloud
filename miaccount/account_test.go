package miaccount

import (
	"testing"
)

func TestGetMicloudCookie(t *testing.T) {
	account := New("")
	micloudCookie, err := account.GenMiCloudCookie()
	if err != nil {
		t.Error(err)
	}
	t.Log(micloudCookie)
}
