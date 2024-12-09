package service

import (
	"context"
	"errors"
	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"

	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
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
	notesResponse := convertToNotesResponse(dbNotes)
	log.Info("Successfully retrieved notes from database")
	return notesResponse, nil
}

func (ns *notesService) CreateNote(ctx context.Context, createNoteRequest dto_model.CreateNoteRequest) error {
	log := logger.Logger()
	log.Info("Checking if note already exists")
	exists, err := ns.repository.NoteExists(ctx, createNoteRequest.Title)
	if err != nil {
		log.Error(err)
		return err
	}
	if exists {
		log.Error("Note already exists")
		return errors.New("note already exists")
	}
	log.Info("Note does not exist. Creating note")
	noteDbModel := convertToNoteEntity(createNoteRequest)
	err = ns.repository.AddNote(ctx, noteDbModel)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("Successfully added note to database")
	return nil
}

func convertToNotesResponse(dbNotes []db_model.Note) dto_model.GetNotesResponse {
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
	return notesResponse
}

func convertToNoteEntity(model dto_model.CreateNoteRequest) db_model.Note {
	noteEntity := db_model.Note{
		Title:       model.Title,
		Description: model.Description,
	}
	return noteEntity
}

func NewNotesService(repository repository.NotesRepository) NotesService {
	return &notesService{
		repository,
	}
}
