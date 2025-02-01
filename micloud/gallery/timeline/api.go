package timeline

import (
	"fmt"
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"time"
)

const (
	timelineApi = "https://i.mi.com/gallery/user/timeline"
)

func GetTimeline(client *client.Client, albumId string) (Timeline, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"albumId", albumId},
	}

	body, r, err := client.DoRequest(request.NewGet(timelineApi, q))
	if err != nil {
		return Timeline{}, err
	}
	return validate.Validate[Timeline](r, body)
}
