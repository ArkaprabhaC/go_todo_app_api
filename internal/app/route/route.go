package route

import (
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/controller"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/repository"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


func InitRoutes(r *gin.Engine, db *sqlx.DB) {

	notesRepository := repository.NewNotesRepository(db)
	notesService := service.NewNotesService(*notesRepository)
	notesController := controller.NewNotesController(*notesService)

	rgV1 := r.Group("/api/v1")
	{
		route := rgV1.Group("/notes")
		{
			route.POST("/", notesController.CreateNote)
		}
	}

}