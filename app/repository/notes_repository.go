package repository

import (
	"context"
	"github.com/jmoiron/sqlx"

	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	_ "github.com/lib/pq"
)

//go:generate mockgen -destination=./mocks/mock_notes_repository.go -package repository_mock -source notes_repository.go
type NotesRepository interface {
	AddNote(ctx context.Context, note db_model.Note) error
}

type notesRepository struct {
	db *sqlx.DB
}

func (n *notesRepository) AddNote(ctx context.Context, note db_model.Note) error {
	tx, err := n.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("INSERT INTO note(title, description) VALUES ($1, $2)", note.Title, note.Description)
	if err != nil {
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
