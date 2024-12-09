package service

import (
	"context"

	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/mapper"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
)

//go:generate mockgen -destination=./mocks/mock_notes_service.go -package service_mock -source notes_service.go
type NotesService interface {
	CreateNote(ctx context.Context, createNoteRequest dto_model.CreateNoteRequest) error
	GetNotes(ctx context.Context) (dto_model.GetNotesResponse, error)
}
type notesService struct {
	repository repository.NotesRepository
}

func (ns *notesService) GetNotes(ctx context.Context) (dto_model.GetNotesResponse, error) {
	log := logger.Logger()
	log.Info("Getting notes from database")
	dbNotes, err := ns.repository.GetNotes(ctx)
	if err != nil {
		log.Error(err)
		return dto_model.GetNotesResponse{}, err
	}
	var notesResponse dto_model.GetNotesResponse
	for _, val := range dbNotes {
		notesResponse.Notes = append(
			notesResponse.Notes,
			dto_model.Note{
				Title:       val.Title,
				Description: val.Description,
			},
		)
	}
	log.Info("Successfully retrieved notes from database")
	return notesResponse, nil
}

func (ns *notesService) CreateNote(ctx context.Context, createNoteRequest dto_model.CreateNoteRequest) error {
	log := logger.Logger()
	log.Info("Incoming create note request - adding note to database ")
	noteDbModel := mapper.ConvertNotesDTOToNotesEntity(createNoteRequest)
	err := ns.repository.AddNote(ctx, noteDbModel)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("Successfully added note to database")
	return nil
}

func NewNotesService(repository repository.NotesRepository) NotesService {
	return &notesService{
		repository,
	}
}
