package main

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/controller"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service"
	"github.com/jmoiron/sqlx"
)

func IntializeNotesController(db *sqlx.DB) controller.NotesController {
	notesRepository := repository.NewNotesRepository(db)
	notesService := service.NewNotesService(notesRepository)
	return controller.NewNotesController(notesService)
}