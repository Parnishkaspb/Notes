package notes

type Note interface {
	CreateNote(id int, value interface{}) (bool, error)
	ReadNote(id int) (interface{}, error)
	UpdateNote(id int, value interface{}) (bool, error)
	DeleteNote(id int) (bool, error)
	ReadAllNotes() (map[int]interface{}, error)
}
