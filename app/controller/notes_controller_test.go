package controller_test

import (
	"bytes"
	"encoding/json"
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
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)


type NotesControllerTestSuite struct {
	suite.Suite
	mockDB *sqlx.DB
	ctrl *gomock.Controller
	engine *gin.Engine
	mockNotesService *service_mock.MockNotesService
	notesController controller.NotesController
}

func (suite *NotesControllerTestSuite) SetupTest() {
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	
	suite.mockDB = sqlx.NewDb(mockDB, "postgres")
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockNotesService = service_mock.NewMockNotesService(suite.ctrl)
	suite.notesController = controller.NewNotesController(suite.mockNotesService)
	suite.engine = gin.Default()
	route.InitializeRoutes(suite.mockDB, suite.engine, suite.notesController)
}

func TestNotesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(NotesControllerTestSuite))
}



func (suite *NotesControllerTestSuite) Test_CreateNote_ShouldAddANoteSuccessfully () {
	
	w := httptest.NewRecorder()
	reqBody := []byte(`{"title": "New Note", "description": "Some note description"}`)
	bodyReader := bytes.NewReader(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/notes/", bodyReader)
	suite.mockNotesService.EXPECT().CreateNote(gomock.Any()).Return(nil)
	
	suite.engine.ServeHTTP(w, req)

	resp := make(map[string]interface{})
	json.Unmarshal(w.Body.Bytes(), &resp)
	expected := "Note created successfully!"
	suite.Equal(201, w.Code)
	suite.Equal(expected, resp["message"])
}

