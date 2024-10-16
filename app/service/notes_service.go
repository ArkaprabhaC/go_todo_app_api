package service

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
)

//go:generate mockgen -destination=./mocks/mock_notes_service.go -package service_mock -source notes_service.go
type NotesService interface {
	CreateNote(createNoteRequest dto_model.CreateNoteRequest) error
}
type notesService struct {
	repository repository.NotesRepository
}

func (ns notesService) CreateNote(createNoteRequest dto_model.CreateNoteRequest) error {
	log := logger.Logger()
	log.Info("Transforming createNoteRequest object into db_model.Note object")
	noteDbModel := db_model.Note{
		Title:       createNoteRequest.Title,
		Description: createNoteRequest.Description,
	}
	err := ns.repository.AddNote(noteDbModel)
	if err != nil {
		log.Error(err)
	}
	return err
}

func NewNotesService(repository repository.NotesRepository) NotesService {
	return notesService{
		repository,
	}
}
