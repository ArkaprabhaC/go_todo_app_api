package repository_test

import (
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

type NotesRepositoryTestSuite struct {
	suite.Suite
	db              *sqlx.DB
	notesRepository repository.NotesRepository
}

func (suite *NotesRepositoryTestSuite) SetupTest() {
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()

	suite.db = sqlx.NewDb(mockDB, "postgres")
	suite.notesRepository = repository.NewNotesRepository(suite.db)
}
