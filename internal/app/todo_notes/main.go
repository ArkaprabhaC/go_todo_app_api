package main

import (
	"fmt"

	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/config"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/database"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/route"
	"github.com/gin-gonic/gin"
)

func main() {

	log := logger.Logger()
	log.Infof("Starting up application...")
	db := database.GetDatabaseConnection()

	router := gin.Default()
	route.InitRoutes(router, db)

	config := config.ReadConfig()
	server_addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	router.Run(server_addr)
}
