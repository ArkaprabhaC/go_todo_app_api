package service

import "github.com/ArkaprabhaC/go_todo_app_api/internal/app/repository"

type NotesService struct {
	repository repository.NotesRepository
}

func NewNotesService(repository repository.NotesRepository) *NotesService {
	return &NotesService{
		repository,
	}
}