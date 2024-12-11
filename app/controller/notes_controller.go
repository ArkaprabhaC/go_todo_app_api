package controller

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	appErrors "github.com/ArkaprabhaC/go_todo_app_api/app/model/dto/errors"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service"
	"github.com/gin-gonic/gin"
	"strings"
)

type NotesController interface {
	CreateNoteHandler(ctx *gin.Context)
	GetNotesHandler(ctx *gin.Context)
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
		ctx.AbortWithStatusJSON(400, appErrors.RequestBodyParseError)
		return
	}
	createNoteRequest.Title = strings.TrimSpace(createNoteRequest.Title)
	err := nc.notesService.CreateNote(ctx, createNoteRequest)
	if err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(500, appErrors.FailureToAddNoteError)
		return
	}
	log.Info("Request exiting..")
	ctx.JSON(200, gin.H{
		"message": "Note created successfully",
	})
}

func (nc *notesController) GetNotesHandler(ctx *gin.Context) {
	log := logger.Logger()
	log.Info("Received request to get all the notes")
	response, err := nc.notesService.GetNotes(ctx)
	if err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(500, appErrors.FailureGetNotes)
		return
	}
	log.Info("Request exiting..")
	ctx.JSON(200, response)
}

func NewNotesController(service service.NotesService) NotesController {
	return &notesController{
		notesService: service,
	}
}
