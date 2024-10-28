package entities

type Note struct {
	Notes map[int]interface{}
}

func NewNote() *Note {
	return &Note{
		Notes: make(map[int]interface{}),
	}
}
