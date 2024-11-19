package mapper

import (
	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	dto_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
)
func ConvertNotesDTOToNotesEntity(model dto_model.CreateNoteRequest) db_model.Note {
	noteEntity := db_model.Note{
		Title:       model.Title,
		Description: model.Description,
	}
	return noteEntity
}

