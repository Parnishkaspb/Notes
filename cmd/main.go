package main

import (
	"Project2/internal/infrastructure/file"
	"Project2/internal/interfaces/controllers/notes"
	"fmt"
)

func main() {
	// Создание сервиса для работы с заметками
	noteService := notes.NewNoteService()

	// Создание новых заметок
	noteService.CreateNote(1, "First Note")
	noteService.CreateNote(2, "Second Note")

	// Чтение всех заметок
	allNotes, err := noteService.ReadAllNotes()
	if err != nil {
		fmt.Println("Error reading notes:", err)
		return
	}
	fmt.Println("All notes:", allNotes.Notes)

	// Создание экземпляра WorkWithFileImpl
	fileHandler := &file.WorkWithFileImpl{}

	// Установка данных заметок в WorkWithFileImpl
	fileHandler.SetNote(allNotes)

	// Сохранение данных в файл
	success, err := fileHandler.SaveData("Text.txt")
	if err != nil {
		fmt.Println("Error saving data:", err)
	} else if success {
		fmt.Println("Data saved successfully to Text.txt")
	}

	// Чтение данных из файла
	success, err = fileHandler.LoadData("Text.txt")
	if err != nil {
		fmt.Println("Error loading data:", err)
	} else if success {
		fmt.Println("Data loaded successfully from Text.txt")
		fmt.Println("Loaded notes:", fileHandler.Note.Notes)
	}
}
