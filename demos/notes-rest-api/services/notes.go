package services

import (
	"context"

	"github.com/steevehook/mdc19/notes-rest-api/models"
)

type NotesRepository interface {
	GetNotes(ctx context.Context, limit int) ([]models.Note, error)
	GetNoteByID(ctx context.Context, id string) (models.Note, error)
	CreateNote(ctx context.Context, req models.NoteReq) error
	UpdateNote(ctx context.Context, req models.NoteReq) error
	DeleteNoteByID(ctx context.Context, id string) error
	DeleteNotes(ctx context.Context) error
}

type NotesService struct {
	notesRepo NotesRepository
}

func NewNotes(repo NotesRepository) NotesService {
	return NotesService{
		notesRepo: repo,
	}
}

func (n NotesService) GetNotes() (int, error) {
	panic("implement me")
}

func (n NotesService) GetNoteByID() (int, error) {
	panic("implement me")
}

func (n NotesService) CreateNote() error {
	panic("implement me")
}

func (n NotesService) UpdateNote() error {
	panic("implement me")
}

func (n NotesService) DeleteNoteByID() error {
	panic("implement me")
}

func (n NotesService) DeleteNotes() error {
	panic("implement me")
}
