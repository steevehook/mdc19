package repositories

import (
	"context"

	"upper.io/db.v3/lib/sqlbuilder"

	"github.com/steevehook/mdc19/notes-rest-api/models"
)

const (
	NotesTableName = "notes"
)

type NotesRepository struct {
	db sqlbuilder.Database
}

func NewNotes(db sqlbuilder.Database) NotesRepository {
	return NotesRepository{
		db: db,
	}
}

func (NotesRepository) GetNotes(ctx context.Context, limit int) ([]models.Note, error) {
	panic("implement me")
}

func (NotesRepository) GetNoteByID(ctx context.Context, id string) (models.Note, error) {
	panic("implement me")
}

func (NotesRepository) CreateNote(ctx context.Context, req models.NoteReq) error {
	panic("implement me")
}

func (NotesRepository) UpdateNote(ctx context.Context, req models.NoteReq) error {
	panic("implement me")
}

func (NotesRepository) DeleteNoteByID(ctx context.Context, id string) error {
	panic("implement me")
}

func (NotesRepository) DeleteNotes(ctx context.Context) error {
	panic("implement me")
}
