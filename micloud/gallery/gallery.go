package gallerymgr

import (
	"errors"
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/gallery/album"
	"github.com/clouderhem/micloud/micloud/gallery/gallery"
	"github.com/clouderhem/micloud/micloud/gallery/timeline"
)

type Gallery struct {
	Client *client.Client
}

func (g *Gallery) ListAlbums(pageNum, pageSize int, shared bool) (album.Albums, error) {
	return album.ListAlbums(g.Client, pageNum, pageSize, shared)
}

func (g *Gallery) ListGalleries(query gallery.GalleriesQuery) (gallery.Galleries, error) {
	return gallery.ListGalleries(g.Client, query)
}

func (g *Gallery) GetGalleryStorageUrl(id string) (string, error) {
	url, err := gallery.GetGalleryStorageUrl(g.Client, id)
	if err != nil {
		return "", err
	}
	if url == "" {
		return "", errors.New("cannot get gallery storage url")
	}
	return url, err
}

func (g *Gallery) GetTimeline(albumId string) (timeline.Timeline, error) {
	return timeline.GetTimeline(g.Client, albumId)
}

func (g *Gallery) DeleteGallery(id string) error {
	return gallery.DeleteGallery(g.Client, id)
}

func (g *Gallery) GetGalleryFile(storageUrl string) (gallery.GalleryFile, error) {
	return gallery.GetGalleryFile(g.Client, storageUrl)
}
