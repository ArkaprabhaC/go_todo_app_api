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
	engine := gin.New()
	notesController := IntializeNotesController(db)
	route.InitializeRoutes(engine, notesController)

	log.Info("Starting application")
	configVars := config.ReadConfig()
	host := configVars.Server.Host
	port := configVars.Server.Port
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	log.Info("======Ready to accept requests======")
	_ = engine.Run(serverAddr)
}
