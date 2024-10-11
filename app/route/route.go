package route

import (

	"github.com/ArkaprabhaC/go_todo_app_api/app/controller"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)


func InitializeRoutes(db *sqlx.DB, engine *gin.Engine, notesController controller.NotesController) {

	rgV1 := engine.Group("/api/v1")
	{
		route := rgV1.Group("/notes")
		{
			route.POST("/", notesController.CreateNoteHandler)
		}
	}

}