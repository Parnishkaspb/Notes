package notes

import (
	"Project2/internal/entities"
	"errors"
)

type NoteImpl struct {
	note *entities.Note
}

func NewNoteService() *NoteImpl {
	return &NoteImpl{
		note: entities.NewNote(),
	}
}

func (n *NoteImpl) CreateNote(id int, value interface{}) (bool, error) {
	if _, err := n.ReadNote(id); err == nil {
		return false, errors.New("note already exists")
	}
	if value == "" {
		return false, errors.New("note value is empty")
	}
	n.note.Notes[id] = value
	return true, nil
}

func (n *NoteImpl) ReadNote(id int) (interface{}, error) {
	value, exists := n.note.Notes[id]
	if !exists {
		return nil, errors.New("note not found")
	}
	return value, nil
}

func (n *NoteImpl) ReadAllNotes() (*entities.Note, error) {
	if len(n.note.Notes) == 0 {
		return nil, errors.New("no notes found")
	}

	return n.note, nil
}
func (n *NoteImpl) UpdateNote(id int, value interface{}) (bool, error) {
	if _, err := n.ReadNote(id); err != nil {
		return false, errors.New("note not found")
	}
	if value == "" {
		return false, errors.New("note value is empty")
	}
	n.note.Notes[id] = value
	return true, nil
}

func (n *NoteImpl) DeleteNote(id int) (bool, error) {
	if _, err := n.ReadNote(id); err != nil {
		return false, errors.New("note not found")
	}
	delete(n.note.Notes, id)
	return true, nil
}
