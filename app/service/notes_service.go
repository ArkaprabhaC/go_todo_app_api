package service

import (
	"context"

	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
	"github.com/ArkaprabhaC/go_todo_app_api/app/mapper"
)

//go:generate mockgen -destination=./mocks/mock_notes_service.go -package service_mock -source notes_service.go
type NotesService interface {
	CreateNote(ctx context.Context, createNoteRequest dto_model.CreateNoteRequest) error
}
type notesService struct {
	repository repository.NotesRepository
}

func (ns *notesService) CreateNote(ctx context.Context, createNoteRequest dto_model.CreateNoteRequest) error {
	log := logger.Logger()
	log.Info("Transforming createNoteRequest object into db_model.Note object")
	noteDbModel := mapper.ConvertNotesDTOToNotesEntity(createNoteRequest)
	err := ns.repository.AddNote(ctx, noteDbModel)
	if err != nil {
		log.Error(err)
	}
	return err
}

func NewNotesService(repository repository.NotesRepository) NotesService {
	return &notesService{
		repository,
	}
}
