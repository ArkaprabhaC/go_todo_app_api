package controller

import (
	"fmt"

	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto/errors"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService service.NotesService
}

func (nc NotesController) CreateNoteHandler(ctx *gin.Context) {
	log := logger.Logger()
	log.Info("Received request to create note")
	var createNoteRequest dto.CreateNoteRequest
	if err := ctx.BindJSON(&createNoteRequest); err != nil {
		log.Errorf("Error while binding request JSON. %v", err)
		ctx.AbortWithStatusJSON(400, errors.REQUEST_BODY_PARSE_ERROR)
	}
	err := nc.notesService.CreateNote(createNoteRequest)
	fmt.Println(err)
	if err != nil {
		ctx.AbortWithStatusJSON(500, errors.FAILURE_TO_ADD_NOTE_ERROR)
	}
	ctx.JSON(201, gin.H{
		"message": "Note created successfully!",
	})

}


func NewNotesController(service service.NotesService) NotesController {
	return NotesController{
		notesService: service,
	}
}