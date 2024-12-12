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

	noteEntity := db_model.Note{
		Title:       createNoteRequest.Title,
		Description: createNoteRequest.Description,
	}

	suite.mockRepository.EXPECT().NoteExists(suite.context, createNoteRequest.Title).Return(false, nil)
	suite.mockRepository.EXPECT().AddNote(suite.context, noteEntity).Return(nil)

	err := suite.service.CreateNote(suite.context, createNoteRequest)
	suite.Nil(err)
}

func (suite *NotesServiceTestSuite) Test_CreateNote_ShouldThrowError_IfNoteWithSameTitleExists() {
	createNoteRequest := dto_model.CreateNoteRequest{
		Title:       "Title",
		Description: "Some description",
	}

	suite.mockRepository.EXPECT().NoteExists(suite.context, createNoteRequest.Title).Return(true, nil)

	err := suite.service.CreateNote(suite.context, createNoteRequest)
	suite.Error(err)
}

func (suite *NotesServiceTestSuite) Test_CreateNote_ShouldThrowError_IfRepositoryErrors_WhileCheckingForNoteExists() {
	createNoteRequest := dto_model.CreateNoteRequest{
		Title:       "Title",
		Description: "Some description",
	}

	suite.mockRepository.EXPECT().NoteExists(suite.context, createNoteRequest.Title).Return(false, errors.New("repository error"))

	err := suite.service.CreateNote(suite.context, createNoteRequest)
	suite.Error(err)
}

func (suite *NotesServiceTestSuite) Test_CreateNote_ShouldThrowError_IfRepositoryErrorsOut() {
	createNoteRequest := dto_model.CreateNoteRequest{
		Title:       "Title",
		Description: "Some description",
	}

	addNote := db_model.Note{
		Title:       createNoteRequest.Title,
		Description: createNoteRequest.Description,
	}

	suite.mockRepository.EXPECT().NoteExists(suite.context, createNoteRequest.Title).Return(false, nil)
	suite.mockRepository.EXPECT().AddNote(suite.context, addNote).Return(errors.New("some repo error occurred"))

	err := suite.service.CreateNote(suite.context, createNoteRequest)
	suite.NotNil(err)
	suite.Equal("Unable to add a note in the system", err.Error())
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

func (suite *NotesServiceTestSuite) Test_GetNotes_ShouldReturnEmptyArray_IfNoNotesAddedInSystem() {
	dbNotes := make([]db_model.Note, 0)
	suite.mockRepository.EXPECT().GetNotes(suite.context).Return(dbNotes, nil)
	notesResponse, err := suite.service.GetNotes(suite.context)
	suite.Nil(err)
	suite.NotNil(notesResponse.Notes)
	suite.Equal(0, len(notesResponse.Notes))
}

func (suite *NotesServiceTestSuite) Test_GetNotes_ShouldThrowError_IfRepositoryErrorsOut() {
	suite.mockRepository.EXPECT().GetNotes(suite.context).Return(nil, errors.New("some repo error occurred"))
	response, err := suite.service.GetNotes(suite.context)
	suite.Error(err)
	suite.Equal(dto_model.GetNotesResponse{}, response)
}
