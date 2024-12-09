package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	dto_model "github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArkaprabhaC/go_todo_app_api/app/controller"
	"github.com/ArkaprabhaC/go_todo_app_api/app/route"
	"github.com/ArkaprabhaC/go_todo_app_api/app/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type NotesControllerTestSuite struct {
	suite.Suite
	mockDB           *sqlx.DB
	ctrl             *gomock.Controller
	engine           *gin.Engine
	mockNotesService *service_mock.MockNotesService
	notesController  controller.NotesController
}

func (suite *NotesControllerTestSuite) SetupTest() {
	suite.mockDB = &sqlx.DB{}
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockNotesService = service_mock.NewMockNotesService(suite.ctrl)
	suite.notesController = controller.NewNotesController(suite.mockNotesService)
	suite.engine = gin.New()
	route.InitializeRoutes(suite.engine, suite.notesController)
}

func TestNotesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(NotesControllerTestSuite))
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldAddANoteSuccessfully() {

	w := httptest.NewRecorder()
	reqBody := []byte(`{"title": "New Note", "description": "Some note description"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.mockNotesService.EXPECT().CreateNote(gomock.Any(), gomock.Any()).Return(nil)

	suite.engine.ServeHTTP(w, req)

	resp := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	expected := "Note created successfully"
	suite.Equal(200, w.Code)
	suite.Equal(expected, resp["message"])
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasMissingDescription() {

	w := httptest.NewRecorder()
	reqBody := []byte(`{"title": "New Note"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(w, req)

	suite.Equal(400, w.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasMissingTitle() {

	w := httptest.NewRecorder()
	reqBody := []byte(`{"name": "New Note"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(w, req)

	suite.Equal(400, w.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorWhenUnableToAddMessage() {
	w := httptest.NewRecorder()
	reqBody := []byte(`{"title": "New Note", "description": "Some note description"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)
	suite.mockNotesService.EXPECT().CreateNote(gomock.Any(), gomock.Any()).Return(errors.New("Some error occurred internally"))

	suite.engine.ServeHTTP(w, req)

	resp := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	expected := "Unable to add a note in the system"
	suite.Equal(500, w.Code)
	suite.Equal(expected, resp["message"])
}

func (suite *NotesControllerTestSuite) Test_GetNotesHandler_ShouldDisplayAllNotes() {
	stubbedResponse := dto_model.GetNotesResponse{
		Notes: []dto_model.Note{
			{
				Title:       "Title 1",
				Description: "Description 1",
			},
			{
				Title:       "Title 2",
				Description: "Description 2",
			},
			{
				Title:       "Title 3",
				Description: "Description 3",
			},
		},
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)
	suite.mockNotesService.EXPECT().GetNotes(gomock.Any()).Return(stubbedResponse, nil)

	suite.engine.ServeHTTP(w, req)

	var actualResponse dto_model.GetNotesResponse
	_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	suite.Equal(200, w.Code)
	suite.Equal(stubbedResponse, actualResponse)

}

func (suite *NotesControllerTestSuite) Test_GetNotesHandler_ShouldReturnError_WhenServiceReturnsError() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)
	suite.mockNotesService.EXPECT().GetNotes(gomock.Any()).Return(dto_model.GetNotesResponse{}, errors.New("some error occurred internally"))

	suite.engine.ServeHTTP(w, req)

	suite.Equal(500, w.Code)
}
