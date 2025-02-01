package gallerymgr

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/gallery/gallery"
	"testing"
)

func TestListAlbums(t *testing.T) {
	g := Gallery{Client: client.GlobalClient}
	albums, err := g.ListAlbums(0, 10, false)
	if err != nil {
		t.Error(err)
	}
	if len(albums.Albums) == 0 {
		t.Error("no albums")
	}
	fmt.Println(albums)
}

func TestListGalleries(t *testing.T) {
	g := Gallery{Client: client.GlobalClient}
	galleries, err := g.ListGalleries(gallery.GalleriesQuery{
		StartDate: 20210226,
		EndDate:   20210226,
		PageNum:   0,
		PageSize:  10,
		AlbumId:   "1",
	})
	if err != nil {
		t.Error(err)
	}
	if len(galleries.Galleries) == 0 {
		t.Error("no galleries")
	}
	fmt.Println(galleries)
}

func TestGetTimeline(t *testing.T) {
	g := Gallery{Client: client.GlobalClient}
	timeline, err := g.GetTimeline("1")
	if err != nil {
		t.Error(err)
	}
	if len(timeline.DayCount) == 0 {
		t.Error("no day count")
	}
	fmt.Println(timeline)
}

func TestDeleteGallery(t *testing.T) {
	g := Gallery{Client: client.GlobalClient}
	err := g.DeleteGallery("1")
	if err != nil {
		t.Error(err)
	}
}

func TestGetGalleryFileUrl(t *testing.T) {
	g := Gallery{Client: client.GlobalClient}
	url, err := g.GetGalleryStorageUrl("1")
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("no url")
	}
	fmt.Println(url)
}
