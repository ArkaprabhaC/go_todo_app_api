package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/config"
	"github.com/ArkaprabhaC/go_todo_app_api/internal/app/logger"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var POSTGRES = "postgres"
var MIGRATION_FILES_PATH = "file://internal/app/database/migrations"
var DATA_SOURCE_URL = createDataSourceString()

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
	db, err := sql.Open(POSTGRES, DATA_SOURCE_URL)
	if err != nil {
		log.Fatalf("Unable to open a connection to database. %s", err)
	}
	
	if err = db.Ping(); err != nil {
		log.Fatalf("Unable to ping database. %s", err)
	}
	return db
}


func RunMigrations() {
	log := logger.Logger()
	m, err := migrate.New(MIGRATION_FILES_PATH, createDataSourceString())
	if err != nil {
		log.Fatalf("Error while running migrations: %s", err)
	}

	if err = m.Up(); err != migrate.ErrNoChange && err != nil {
		log.Fatalf("Error while running migrations: %s", err)
	}
}