package repository_test

import (
	"context"
	"errors"
	"testing"

	db_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

type NotesRepositoryTestSuite struct {
	suite.Suite
	context         context.Context
	sqlMock         sqlmock.Sqlmock
	mockDb          *sqlx.DB
	notesRepository repository.NotesRepository
}

func TestNotesRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(NotesRepositoryTestSuite))
}

func (suite *NotesRepositoryTestSuite) SetupSuite() {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	suite.context = context.TODO()
	suite.mockDb = db
	suite.sqlMock = mock
	suite.notesRepository = repository.NewNotesRepository(db)
}

func (suite *NotesRepositoryTestSuite) TearDownSuite() {
	suite.mockDb.Close()
}

func (suite *NotesRepositoryTestSuite) Test_AddNote_ShouldAddNoteInDbSuccessfully() {
	note := db_model.Note{
		Title:       "Note 1",
		Description: "Note description 1",
	}

	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(`INSERT INTO`).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlMock.ExpectCommit()

	err := suite.notesRepository.AddNote(suite.context, note)
	suite.Nil(err)
	suite.Nil(suite.sqlMock.ExpectationsWereMet())

}

func (suite *NotesRepositoryTestSuite) Test_AddNote_ShouldThrowErrorIfQueryExecutionFails() {
	note := db_model.Note{
		Title:       "Note 1",
		Description: "Note description 1",
	}

	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(`INSERT INTO`).WillReturnError(errors.New("some error occurred"))

	err := suite.notesRepository.AddNote(suite.context, note)
	suite.Error(err)
	suite.Nil(suite.sqlMock.ExpectationsWereMet())
}

func (suite *NotesRepositoryTestSuite) Test_AddNote_ShouldThrowErrorIfTransactionFailsToCommit() {
	note := db_model.Note{
		Title:       "Note 1",
		Description: "Note description 1",
	}

	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectExec(`INSERT INTO`).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlMock.ExpectCommit().WillReturnError(errors.New("err occurred"))

	err := suite.notesRepository.AddNote(suite.context, note)
	suite.Error(err)
	suite.Nil(suite.sqlMock.ExpectationsWereMet())
}

func (suite *NotesRepositoryTestSuite) Test_GetNotes_ShouldReturnNotesFromDBSuccessfully() {

	suite.sqlMock.ExpectQuery("SELECT title, description FROM note").
		WillReturnRows(
			sqlmock.NewRows([]string{"title", "description"}).
				AddRow("Note 1", "Note description 1"),
		)

	notes, err := suite.notesRepository.GetNotes(suite.context)
	suite.Nil(err)
	suite.Nil(suite.sqlMock.ExpectationsWereMet())
	suite.Equal(1, len(notes))
	suite.Equal(notes[0].Title, "Note 1")
	suite.Equal(notes[0].Description, "Note description 1")
}

func (suite *NotesRepositoryTestSuite) Test_GetNotes_ShouldReturnError_IfQueryExecutionFails() {
	suite.sqlMock.ExpectQuery("SELECT title, description FROM note").
		WillReturnError(errors.New("some error occurred"))

	_, err := suite.notesRepository.GetNotes(suite.context)
	suite.Nil(suite.sqlMock.ExpectationsWereMet())
	suite.Error(err)
}
