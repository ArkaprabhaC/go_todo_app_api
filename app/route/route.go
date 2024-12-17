package route

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/controller"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(engine *gin.Engine, notesController controller.NotesController) {

	rgv1 := engine.Group("/api/v1")
	{
		route := rgv1.Group("/notes")
		{
			route.POST("", notesController.CreateNoteHandler)
			route.GET("", notesController.GetNotesHandler)
			route.DELETE("/:noteTitle", notesController.DeleteNoteHandler)

		}
	}

}
