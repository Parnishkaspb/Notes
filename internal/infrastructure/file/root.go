package file

import (
	"Project2/internal/entities"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type WorkWithFileImpl struct {
	Note *entities.Note
}

func (file *WorkWithFileImpl) SetNote(note *entities.Note) {
	file.Note = note
}

func (file *WorkWithFileImpl) SaveData(filename string) (bool, error) {
	if file.Note == nil || len(file.Note.Notes) == 0 {
		return false, errors.New("no data to save")
	}

	f, err := os.Create(filename)
	if err != nil {
		return false, errors.New("can't create new file: " + err.Error())
	}
	defer f.Close()

	var stringData []string
	for _, v := range file.Note.Notes {
		if str, ok := v.(string); ok {
			stringData = append(stringData, str)
		}
	}

	_, err = f.WriteString(strings.Join(stringData, "\n"))
	if err != nil {
		return false, errors.New("can't write to file: " + err.Error())
	}

	fmt.Println("Data successfully saved to:", filename)
	return true, nil
}

func (file *WorkWithFileImpl) LoadData(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, errors.New("can't open file: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	notes := make(map[int]interface{})
	index := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			notes[index] = line
			index++
		}
	}

	if err := scanner.Err(); err != nil {
		return false, errors.New("error reading file: " + err.Error())
	}

	file.Note = &entities.Note{Notes: notes}

	fmt.Println("Data successfully loaded from:", filename)
	return true, nil
}
