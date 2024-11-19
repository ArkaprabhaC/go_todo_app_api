package controller

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto/errors"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service"
	"github.com/gin-gonic/gin"
)

type NotesController interface {
	CreateNoteHandler(ctx *gin.Context)
}
type notesController struct {
	notesService service.NotesService
}

func (nc *notesController) CreateNoteHandler(ctx *gin.Context) {
	log := logger.Logger()
	log.Info("Received request to create note")
	var createNoteRequest dto_model.CreateNoteRequest
	if err := ctx.BindJSON(&createNoteRequest); err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(400, errors.REQUEST_BODY_PARSE_ERROR)
		return
	}
	err := nc.notesService.CreateNote(ctx, createNoteRequest)
	if err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(500, errors.FAILURE_TO_ADD_NOTE_ERROR)
		return
	}
	ctx.JSON(201, gin.H{
		"message": "Note created successfully!",
	})

}

func NewNotesController(service service.NotesService) NotesController {
	return &notesController{
		notesService: service,
	}
}
