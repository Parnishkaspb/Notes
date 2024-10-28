package file

import "Project2/internal/entities"

type WorkWithFile interface {
	LoadData(filename string) (entities.Note, error)
	SaveData(filename string, note entities.Note) (bool, error)
	doesFileExist(filename string) (bool, error)
}
