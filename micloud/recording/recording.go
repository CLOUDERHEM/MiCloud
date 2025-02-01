package recordingmgr

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/recording/recording"
)

type Recording struct {
	Client *client.Client
}

func (r *Recording) ListRecordings(offset, limit int) ([]recording.Recording, error) {
	return recording.ListRecordings(r.Client, offset, limit)
}

func (r *Recording) GetRecordingFileUrl(id string) (string, error) {
	return recording.GetRecordingFileUrl(r.Client, id)
}

func (r *Recording) DeleteRecording(id string) error {
	return recording.DeleteRecording(r.Client, id)
}
