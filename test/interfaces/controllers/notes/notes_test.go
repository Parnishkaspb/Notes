package notes

import (
	"Project2/internal/interfaces/controllers/notes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWorkWithFile(t *testing.T) {
	noteService := notes.NewNoteService()

	t.Run("Create Note", func(t *testing.T) {
		ok, err := noteService.CreateNote(1, "Hello, world!")
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("Read Note", func(t *testing.T) {
		note, err := noteService.ReadNote(1)
		require.NoError(t, err)
		assert.Equal(t, "Hello, world!", note)
	})

	t.Run("Read All Notes", func(t *testing.T) {
		notes, err := noteService.ReadAllNotes()
		require.NoError(t, err)

		expectedNotes := map[int]any{
			1: "Hello, world!",
		}

		assert.Equal(t, expectedNotes, notes.Note)
	})

	t.Run("Update Note", func(t *testing.T) {
		ok, err := noteService.UpdateNote(1, "Hello, Any!")
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("Delete Note", func(t *testing.T) {
		ok, err := noteService.DeleteNote(1)
		require.NoError(t, err)
		assert.True(t, ok)
	})
}
