package controller

import (
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/service"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService service.NotesService
}

func (nc NotesController) CreateNote(ctx *gin.Context) {
	log := logger.Logger()
	log.Infof("Starting up application...")
}


func NewNotesController(service service.NotesService) *NotesController {
	return &NotesController{
		notesService: service,
	}
}