package services

func NewNotes() NotesService {
	return NotesService{}
}

type NotesService struct {
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

func (n NotesService) DeleteNote() error {
	panic("implement me")
}
