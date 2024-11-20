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
	log.Infof("Getting db connection and running migration.")
	
	database.RunMigrations()
	db := database.GetDatabaseConnection()

	log.Info("Initializing engine, routes and dependencies")
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	notes_controller := IntializeNotesController(db)
	route.InitializeRoutes(db, engine, notes_controller)

	log.Info("Starting application")
	config := config.ReadConfig()
	host := config.Server.Host
	port := config.Server.Port
	server_addr := fmt.Sprintf("%s:%d", host, port)
	log.Info("======Ready to accept requests======")
	engine.Run(server_addr)
}
