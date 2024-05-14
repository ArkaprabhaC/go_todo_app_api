package database

import (
	"database/sql"
	"fmt"

	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/config"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/todo_notes/logger"
	_ "github.com/lib/pq"
)

func createDataSourceString() (datasource string) {
	const postgresDatasource = "postgres://%s:%s@%s:%s/%s?sslmode=%s"

	dbConfig := config.ReadConfig().Database
	datasource = fmt.Sprintf(postgresDatasource,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.HostName,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.SslMode,
	)
	return
}

func GetDatabaseConnection() *sql.DB {
	log := logger.Logger()
	db, err := sql.Open("postgres", createDataSourceString())
	if err != nil {
		log.Fatalf("Unable to open a connection to database. ", err)
	}
	
	if err = db.Ping(); err != nil {
		log.Fatalf("Unable to ping database.", err)
	}
	return db
}
