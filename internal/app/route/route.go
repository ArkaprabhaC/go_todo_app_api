package route

import (
	"time"

	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/controller"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/middleware"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/model/dto/errors"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/repository"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


func InitializeRoutes(engine *gin.Engine, db *sqlx.DB) {

	notesRepository := repository.NewNotesRepository(db)
	notesService := service.NewNotesService(*notesRepository)
	notesController := controller.NewNotesController(*notesService)

	engine.Use(middleware.RequestTimeout(3*time.Second, errors.NewRequestTimeoutError()))
	rgV1 := engine.Group("/api/v1")
	{
		route := rgV1.Group("/notes")
		{
			route.GET("/", notesController.CreateNote)
		}
	}

}