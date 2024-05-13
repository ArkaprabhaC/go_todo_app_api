package route

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {

	apiV1 := r.Group("/api/v1/")
	{
		apiV1.GET("/hello", func(ctx *gin.Context) { ctx.JSON(200, "Hello World")})
	}

}