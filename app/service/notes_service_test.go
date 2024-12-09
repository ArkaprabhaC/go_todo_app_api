package service_test

import (
	"context"
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
	context        context.Context
	mockRepository *repository_mock.MockNotesRepository
	service        service.NotesService
}

func TestNotesServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NotesServiceTestSuite))
}

func (suite *NotesServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.context = context.TODO()
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

	suite.mockRepository.EXPECT().AddNote(suite.context, addNote).Return(nil)

	err := suite.service.CreateNote(suite.context, createNoteRequest)
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

	suite.mockRepository.EXPECT().AddNote(suite.context, addNote).Return(errors.New("Some repo error occurred"))

	err := suite.service.CreateNote(suite.context, createNoteRequest)
	suite.NotNil(err)
	suite.Equal("Some repo error occurred", err.Error())
}

func (suite *NotesServiceTestSuite) Test_GetNotes_ShouldGetNotesSuccessfully() {
	dbNotes := []db_model.Note{
		{
			Title:       "Title 1",
			Description: "Description 1",
		},
		{
			Title:       "Title 2",
			Description: "Description 2",
		},
	}
	suite.mockRepository.EXPECT().GetNotes(suite.context).Return(dbNotes, nil)
	notesResponse, err := suite.service.GetNotes(suite.context)
	suite.Nil(err)
	suite.Equal(2, len(notesResponse.Notes))
}

func (suite *NotesServiceTestSuite) Test_GetNotes_ShouldThrowErrorIfRepositoryErrorsOut() {
	suite.mockRepository.EXPECT().GetNotes(suite.context).Return(nil, errors.New("some repo error occurred"))
	response, err := suite.service.GetNotes(suite.context)
	suite.Error(err)
	suite.Equal(dto_model.GetNotesResponse{}, response)
}
