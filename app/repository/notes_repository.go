package repository

import (
	"context"
	"github.com/jmoiron/sqlx"

	dbmodel "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	_ "github.com/lib/pq"
)

//go:generate mockgen -destination=./mocks/mock_notes_repository.go -package repository_mock -source notes_repository.go
type NotesRepository interface {
	AddNote(ctx context.Context, note dbmodel.Note) error
	GetNotes(ctx context.Context) ([]dbmodel.Note, error)
	NoteExists(ctx context.Context, title string) (bool, error)
	DeleteNote(ctx context.Context, title string) error
}

const (
	createNoteQuery = "INSERT INTO note(title, description) VALUES ($1, $2)"
	getAllNoteQuery = "SELECT title, description FROM note"
	checkNoteExists = "SELECT EXISTS(SELECT 1 FROM note WHERE title = $1)"
	deleteNoteQuery = "DELETE FROM note WHERE title = $1"
)

type notesRepository struct {
	db *sqlx.DB
}

func (n *notesRepository) DeleteNote(ctx context.Context, title string) error {
	tx, err := n.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.Exec(deleteNoteQuery, title)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil

}

func (n *notesRepository) NoteExists(ctx context.Context, title string) (bool, error) {
	var noteExists bool
	err := n.db.GetContext(ctx, &noteExists, checkNoteExists, title)
	if err != nil {
		return false, err
	}
	return noteExists, nil
}

func (n *notesRepository) GetNotes(ctx context.Context) ([]dbmodel.Note, error) {
	var notes []dbmodel.Note
	err := n.db.SelectContext(ctx, &notes, getAllNoteQuery)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *notesRepository) AddNote(ctx context.Context, note dbmodel.Note) error {
	tx, err := n.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tx.Exec(createNoteQuery, note.Title, note.Description)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewNotesRepository(db *sqlx.DB) NotesRepository {
	return &notesRepository{
		db,
	}
}
