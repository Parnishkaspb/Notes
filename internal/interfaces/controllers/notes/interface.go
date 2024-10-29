package notes

type Note interface {
	CreateNote(id int, value any) (bool, error)
	ReadNote(id int) (any, error)
	UpdateNote(id int, value any) (bool, error)
	DeleteNote(id int) (bool, error)
	ReadAllNotes() (map[int]any, error)
}
