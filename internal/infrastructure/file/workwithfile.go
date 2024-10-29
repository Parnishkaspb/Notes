package file

import (
	"Project2/internal/entities"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WorkWithFileImpl struct {
}

func NewWorkWithFileImpl() *WorkWithFileImpl {
	return &WorkWithFileImpl{}
}
func (file *WorkWithFileImpl) SaveData(filename string, notes entities.Notes) (bool, error) {
	if len(notes.Note) == 0 {
		return false, errors.New("no data to save")
	}

	f, err := os.Create(filename)
	if err != nil {
		return false, errors.New("can't create new file: " + err.Error())
	}
	defer f.Close()

	var stringData []string
	for k, v := range notes.Note {
		keyStr := strconv.Itoa(k)
		if valueStr, ok := v.(string); ok {
			stringData = append(stringData, fmt.Sprintf("%s=%s", keyStr, valueStr))
		} else {
			stringData = append(stringData, fmt.Sprintf("%s=%v", keyStr, v))
		}
	}

	_, err = f.WriteString(strings.Join(stringData, "\n"))
	if err != nil {
		return false, errors.New("can't write to file: " + err.Error())
	}

	fmt.Println("Data successfully saved to:", filename)
	return true, nil
}

func (file *WorkWithFileImpl) LoadData(filename string) (entities.Notes, error) {
	f, err := os.Open(filename)
	if err != nil {
		return entities.Notes{}, errors.New("can't open file: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	notes := make(map[int]any)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				return entities.Notes{}, errors.New("invalid line format, expected 'index=value'")
			}

			index, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				return entities.Notes{}, fmt.Errorf("invalid index in line: %s", line)
			}

			value := strings.TrimSpace(parts[1])
			notes[index] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return entities.Notes{}, errors.New("error reading file: " + err.Error())
	}

	fmt.Println("Data successfully loaded from:", filename)
	return entities.Notes{Note: notes}, nil
}
