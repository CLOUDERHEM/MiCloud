package album

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"strconv"
	"time"
)

const (
	listAlbumApi = "https://i.mi.com/gallery/user/album/list"
)

func ListAlbums(client *client.Client, pageNum, pageSize int, shared bool) (Albums, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"_dc", ts},
		{"pageNum", strconv.Itoa(pageNum)},
		{"pageSize", strconv.Itoa(pageSize)},
		{"isShared", strconv.FormatBool(shared)},
		{"numOfThumbnails", "1"},
	}

	body, r, err := client.DoRequest(request.NewGet(listAlbumApi, q))
	if err != nil {
		return Albums{}, err
	}
	return validate.Validate[Albums](r, body)
}
