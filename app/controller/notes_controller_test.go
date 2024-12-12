package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	dtoModel "github.com/ArkaprabhaC/go_todo_app_api/app/model/dto"
	appErrors "github.com/ArkaprabhaC/go_todo_app_api/app/model/dto/errors"
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
	recorder         *httptest.ResponseRecorder
	context          *gin.Context
	mockNotesService *service_mock.MockNotesService
	notesController  controller.NotesController
}

func (suite *NotesControllerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.engine = gin.New()
	suite.mockDB = &sqlx.DB{}
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockNotesService = service_mock.NewMockNotesService(suite.ctrl)
	suite.notesController = controller.NewNotesController(suite.mockNotesService)
	route.InitializeRoutes(suite.engine, suite.notesController)
}

func TestNotesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(NotesControllerTestSuite))
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldAddANoteSuccessfully() {
	createNoteReq := dtoModel.CreateNoteRequest{
		Title:       "New Note",
		Description: "Some note description",
	}
	reqBody := []byte(`{"title": "New Note", "description": "Some note description"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.mockNotesService.EXPECT().CreateNote(gomock.Any(), createNoteReq).Return(nil)

	suite.engine.ServeHTTP(suite.recorder, req)

	resp := make(map[string]interface{})
	_ = json.Unmarshal(suite.recorder.Body.Bytes(), &resp)
	expected := "Note created successfully"
	suite.Equal(200, suite.recorder.Code)
	suite.Equal(expected, resp["message"])
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldAddANoteSuccessfully_AfterCleaningInputFields() {
	createNoteReq := dtoModel.CreateNoteRequest{
		Title:       "New Note",
		Description: "Some note description",
	}
	reqBody := []byte(`{"title": "   New Note        ", "description": "Some note description"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.mockNotesService.EXPECT().CreateNote(gomock.Any(), createNoteReq).Return(nil)

	suite.engine.ServeHTTP(suite.recorder, req)

	resp := make(map[string]interface{})
	_ = json.Unmarshal(suite.recorder.Body.Bytes(), &resp)
	expected := "Note created successfully"
	suite.Equal(200, suite.recorder.Code)
	suite.Equal(expected, resp["message"])
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasMissingDescription() {

	reqBody := []byte(`{"title": "New Note"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(suite.recorder, req)

	suite.Equal(400, suite.recorder.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasMissingTitle() {

	reqBody := []byte(`{"name": "New Note"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(suite.recorder, req)

	suite.Equal(400, suite.recorder.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasEmptyTitle() {

	reqBody := []byte(`{"title": ""}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(suite.recorder, req)

	suite.Equal(400, suite.recorder.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasNoDescription() {

	reqBody := []byte(`{"title": "New Note"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(suite.recorder, req)

	suite.Equal(400, suite.recorder.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorIfPayloadHasEmptyDescription() {

	reqBody := []byte(`{"title": "New Note", "description": ""}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)

	suite.engine.ServeHTTP(suite.recorder, req)

	suite.Equal(400, suite.recorder.Code)
}

func (suite *NotesControllerTestSuite) Test_CreateNoteHandler_ShouldThrowErrorWhenUnableToAddMessage() {
	reqBody := []byte(`{"title": "New Note", "description": "Some note description"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes", bodyReader)
	suite.mockNotesService.EXPECT().CreateNote(gomock.Any(), gomock.Any()).Return(&appErrors.FailureAddNoteError)

	suite.engine.ServeHTTP(suite.recorder, req)

	resp := make(map[string]interface{})
	_ = json.Unmarshal(suite.recorder.Body.Bytes(), &resp)
	expected := "Unable to add a note in the system"
	suite.Equal(500, suite.recorder.Code)
	suite.Equal(expected, resp["message"])
}

func (suite *NotesControllerTestSuite) Test_GetNotesHandler_ShouldDisplayAllNotes() {
	stubbedResponse := dtoModel.GetNotesResponse{
		Notes: []dtoModel.Note{
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
	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)
	suite.mockNotesService.EXPECT().GetNotes(gomock.Any()).Return(stubbedResponse, nil)

	suite.engine.ServeHTTP(suite.recorder, req)

	var actualResponse dtoModel.GetNotesResponse
	_ = json.Unmarshal(suite.recorder.Body.Bytes(), &actualResponse)
	suite.Equal(200, suite.recorder.Code)
	suite.Equal(stubbedResponse, actualResponse)

}

func (suite *NotesControllerTestSuite) Test_GetNotesHandler_ShouldReturnError_WhenServiceReturnsError() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/notes", nil)
	suite.mockNotesService.EXPECT().GetNotes(gomock.Any()).Return(dtoModel.GetNotesResponse{}, errors.New("some error occurred internally"))

	suite.engine.ServeHTTP(w, req)

	suite.Equal(500, w.Code)
}
