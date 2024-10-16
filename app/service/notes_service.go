package service

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
)

//go:generate mockgen -destination=./mocks/mock_notes_service.go -package service_mock -source notes_service.go
type NotesService interface {
	CreateNote(createNoteRequest dto.CreateNoteRequest) error
}
type notesService struct {
	repository repository.NotesRepository
}

func (ns notesService) CreateNote(createNoteRequest dto.CreateNoteRequest) error {
	return nil
}

func NewNotesService(repository repository.NotesRepository) NotesService {
	return notesService{
		repository,
	}
}
