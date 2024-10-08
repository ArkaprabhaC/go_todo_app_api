package main

import (
	"fmt"

	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/config"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/database"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/logger"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/route"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.Logger()
	log.Infof("Starting up application...")
	
	database.RunMigrations()
	db := database.GetDatabaseConnection()

	engine := gin.Default()
	route.InitializeRoutes(engine, db)

	config := config.ReadConfig()
	server_addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	engine.Run(server_addr)

}
