package controller

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	appErrors "github.com/ArkaprabhaC/go_todo_app_api/app/model/errors"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type NotesController interface {
	CreateNoteHandler(ctx *gin.Context)
	GetNotesHandler(ctx *gin.Context)
	DeleteNoteHandler(ctx *gin.Context)
}
type notesController struct {
	notesService service.NotesService
}

func (nc *notesController) DeleteNoteHandler(ctx *gin.Context) {
	log := logger.Logger()
	log.Info("Received request to delete note")
	noteIdString := ctx.Param("id")
	noteId, err := strconv.Atoi(noteIdString)
	if err != nil {
		log.Error("Failed to convert id to int")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, appErrors.FailureBadRequest)
		return
	}
	err = nc.notesService.DeleteNote(ctx, noteId)
	if err != nil {
		log.Error("Failed to delete note")
		ctx.AbortWithStatusJSON(http.StatusNotFound, appErrors.FailureNoteNotFound)
		return
	}
	log.Info("Request exiting..")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Note with ID 123 deleted",
	})
}

func (nc *notesController) CreateNoteHandler(ctx *gin.Context) {
	log := logger.Logger()
	log.Info("Received request to create note")
	var createNoteRequest dto_model.CreateNoteRequest
	if err := ctx.BindJSON(&createNoteRequest); err != nil {
		log.Error("Failed to bind request body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, appErrors.FailureBadRequest)
		return
	}
	createNoteRequest.Title = strings.TrimSpace(createNoteRequest.Title)
	err := nc.notesService.CreateNote(ctx, createNoteRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
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
