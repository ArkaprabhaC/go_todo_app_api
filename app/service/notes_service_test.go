package service_test

import (
	"errors"
	"testing"

	"github.com/ArkaprabhaC/go_todo_app_api/app/model/db"
	"github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"github.com/ArkaprabhaC/go_todo_app_api/app/repository/mocks"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type NotesServiceTestSuite struct {
	suite.Suite
	mockRepository *repository_mock.MockNotesRepository
	service        service.NotesService
}

func TestNotesServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NotesServiceTestSuite))
}

func (suite *NotesServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mockRepository = repository_mock.NewMockNotesRepository(ctrl)
	suite.service = service.NewNotesService(suite.mockRepository)
}

func (suite *NotesServiceTestSuite) Test_CreateNote_ShouldAddNoteSuccessfully() {
	createNoteRequest := dto_model.CreateNoteRequest{
		Title:       "Title",
		Description: "Some description",
	}

	addNote := db_model.Note{
		Title:       createNoteRequest.Title,
		Description: createNoteRequest.Description,
	}

	suite.mockRepository.EXPECT().AddNote(addNote).Return(nil)

	err := suite.service.CreateNote(createNoteRequest)
	suite.Nil(err)
}

func (suite *NotesServiceTestSuite) Test_CreateNote_ShouldThrowErrorIfUnableToAddNote() {
	createNoteRequest := dto_model.CreateNoteRequest{
		Title:       "Title",
		Description: "Some description",
	}

	addNote := db_model.Note{
		Title:       createNoteRequest.Title,
		Description: createNoteRequest.Description,
	}

	suite.mockRepository.EXPECT().AddNote(addNote).Return(errors.New("Some repo error occurred"))

	err := suite.service.CreateNote(createNoteRequest)
	suite.NotNil(err)
	suite.Equal("Some repo error occurred", err.Error())
}
