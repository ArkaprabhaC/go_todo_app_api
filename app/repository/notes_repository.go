package repository

import (
	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination=./mocks/mock_notes_repository.go -package repository_mock -source notes_repository.go
type NotesRepository interface {
	AddNote(note db_model.Note) error
}

type notesRepository struct {
	db *sqlx.DB
}

func (r notesRepository) AddNote(note db_model.Note) error {
	return nil
}

func NewNotesRepository(db *sqlx.DB) NotesRepository {
	return notesRepository{
		db,
	}
}
