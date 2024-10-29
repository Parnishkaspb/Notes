package file

import (
	"Project2/internal/infrastructure/file"
	"Project2/internal/interfaces/controllers/notes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWorkWithFile(t *testing.T) {
	noteService := notes.NewNoteService()
	noteService.CreateNote(1, "First Note")
	noteService.CreateNote(2, "Second Note")
	allNotes, _ := noteService.ReadAllNotes()
	fileHandler := file.NewWorkWithFileImpl()

	t.Run("SaveFile", func(t *testing.T) {
		ok, err := fileHandler.SaveData("Text.txt", allNotes)
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("ReadFile", func(t *testing.T) {
		loadNotes, err := fileHandler.LoadData("Text.txt")
		require.NoError(t, err)
		assert.Equal(t, allNotes, loadNotes)
	})
}
