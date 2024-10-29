package main

import (
	"Project2/internal/infrastructure/file"
	"Project2/internal/interfaces/controllers/notes"
	"log"
)

func main() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime)

	noteService := notes.NewNoteService()
	noteService.CreateNote(1, "First Note")
	noteService.CreateNote(2, "Second Note")

	allNotes, err := noteService.ReadAllNotes()
	if err != nil {
		log.Fatal("ошибка при прочтении всех заметок")
	}
	log.Println("All notes:", allNotes.Note)

	fileHandler := file.NewWorkWithFileImpl()

	if _, err := fileHandler.SaveData("Text.txt", allNotes); err != nil {
		log.Fatal("ошибка при записи в файл")
	}

	loadNotes, err := fileHandler.LoadData("Text.txt")
	if err != nil {
		log.Fatal("невозможно прочитать файл")
	}
	log.Println("Loaded notes:", loadNotes)
}
