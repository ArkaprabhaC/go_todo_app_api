package repository

import (
	"context"
	"errors"
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

func (n notesRepository) AddNote(ctx context.Context, note db_model.Note) error {
	tx, err := n.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("unable to start transaction")
	}

	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO note(title, description) VALUES ($1, $2)", note.Title, note.Description)
	if err != nil {
		return errors.New("unable to insert note to the database")
	}

	if err = tx.Commit(); err != nil {
        return errors.New("unable to commit transaction")
    }

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1{
		return errors.New("unable to get rows affected or invalid rows affected after inserting note")
	}
	return err
}

func NewNotesRepository(db *sqlx.DB) NotesRepository {
	return notesRepository{
		db,
	}
}
