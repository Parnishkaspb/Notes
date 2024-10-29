package notes

import (
	"Project2/internal/entities"
	"errors"
)

type NoteImpl struct {
	note entities.Notes
}

func NewNoteService() *NoteImpl {
	return &NoteImpl{
		note: entities.Notes{make(map[int]any)},
	}
}

func (n *NoteImpl) CreateNote(id int, value any) (bool, error) {
	if value == nil {
		return false, errors.New("note value is empty")
	}

	if _, err := n.ReadNote(id); err == nil {
		return false, errors.New("note already exists")
	}

	n.note.Note[id] = value
	return true, nil
}

func (n *NoteImpl) ReadNote(id int) (any, error) {
	value, exists := n.note.Note[id]
	if !exists {
		return nil, errors.New("note not found")
	}
	return value, nil
}

func (n *NoteImpl) ReadAllNotes() (entities.Notes, error) {
	if len(n.note.Note) == 0 {
		return entities.Notes{}, errors.New("no notes found")
	}

	return entities.Notes{n.note.Note}, nil
}
func (n *NoteImpl) UpdateNote(id int, value any) (bool, error) {
	if value == nil {
		return false, errors.New("note value is empty")
	}

	if _, err := n.ReadNote(id); err != nil {
		return false, errors.New("note not found")
	}

	n.note.Note[id] = value
	return true, nil
}

func (n *NoteImpl) DeleteNote(id int) (bool, error) {
	if _, err := n.ReadNote(id); err != nil {
		return false, errors.New("note not found")
	}
	delete(n.note.Note, id)
	return true, nil
}
