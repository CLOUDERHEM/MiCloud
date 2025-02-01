package notemgr

import (
	"github.com/clouderhem/micloud/client"
	"github.com/clouderhem/micloud/micloud/note/note"
	"github.com/clouderhem/micloud/micloud/note/recyclebin"
	"github.com/clouderhem/micloud/utility/parallel"
	"math/rand"
	"time"
)

type Note struct {
	Client        *client.Client
	NumOfReqInSec int
}

func (n *Note) ListNotes(limit int) (note.Notes, error) {
	return note.ListNotes(n.Client, limit)
}

func (n *Note) GetNote(id string) (note.Note, error) {
	return note.GetNote(n.Client, id)
}

func (n *Note) ListFullNotes(noteIds []string) ([]note.Note, []parallel.ErrOut[string]) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res, errs := parallel.DoParallel[string, note.Note](noteIds,
		func(id string) (note.Note, error) {
			time.Sleep(time.Second *
				time.Duration((r.Intn((len(noteIds)/n.NumOfReqInSec)+1))+1))
			fullNote, err := note.GetNote(n.Client, id)
			return fullNote, err
		})
	return res, errs
}

func (n *Note) DeleteNote(id, tag string, purge bool) error {
	return note.DeleteNote(n.Client, id, tag, purge)
}

func (n *Note) ListDeletedNotes(syncTag string, limit int) (note.Notes, error) {
	return recyclebin.ListDeletedNotes(n.Client, syncTag, limit)
}

// GetNoteFileUrl get file url in note, pic or record
func (n *Note) GetNoteFileUrl(fileId string) (string, error) {
	return note.GetNoteFileUrl(n.Client, note.FileType, fileId)
}
