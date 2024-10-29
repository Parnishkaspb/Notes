package file

import "Project2/internal/entities"

type WorkWithFile interface {
	LoadData(filename string, notes entities.Notes) (bool, error)
	SaveData(filename string) (entities.Notes, error)
}
