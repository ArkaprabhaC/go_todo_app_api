package main

import (
	"fmt"

	"github.com/ArkaprabhaC/go_todo_app_api/app/config"
	"github.com/ArkaprabhaC/go_todo_app_api/app/database"
	"github.com/ArkaprabhaC/go_todo_app_api/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/app/route"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.Logger()
	log.Infof("Starting up application...")
	
	database.RunMigrations()
	db := database.GetDatabaseConnection()

	engine := gin.Default()
	notes_controller := IntializeNotesController(db)
	route.InitializeRoutes(db, engine, notes_controller)

	config := config.ReadConfig()
	host := config.Server.Host
	port := config.Server.Port
	server_addr := fmt.Sprintf("%s:%d", host, port)
	engine.Run(server_addr)

}
