// Boilerplate test file leveraging testify.
// Creates a suite with initial setup for the server, and then executes defined API tests.
package tests

import (
	"erik-craigo-code-challenge/config"
	"erik-craigo-code-challenge/controllers"
	database "erik-craigo-code-challenge/db"
	"erik-craigo-code-challenge/server"
	"erik-craigo-code-challenge/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

// suite struct defines necessary components to execute tests
type StudentSuite struct {
	suite.Suite
	Config   *viper.Viper
	Router   *gin.Engine
	Response *httptest.ResponseRecorder
}

// sets up suite with dev configs and starts sets up the router
// SetupSuite runs once before executing any tests
func (suite *StudentSuite) SetupSuite() {
	config.Init("dev")
	suite.Config = config.GetConfig()
	// set up DBs
	DB := database.Init()

	studentController := controllers.NewStudentController(DB)
	examController := controllers.NewExamController(DB)
	studentExamController := controllers.NewStudentExamController(DB)
	suite.Router = server.SetupRouter(studentController, studentExamController, examController)
	studentDB := database.Student{DB}
	studentExamDB := database.StudentExam{DB}
	// get student test data
	students := tests.GetTestStudents()
	studentExams := tests.GetTestStudentExams()
	// populate DBs
	for _, student := range students {
		studentDB.Insert(student)
	}
	for _, studentExam := range studentExams {
		studentExamDB.Insert(studentExam)
	}

}

// sets up env for each test
// SetupTest runes before each test (after SetupSuite)
func (suite *StudentSuite) SetupTest() {
	suite.Response = httptest.NewRecorder()
}

// test to verify /students functionality
func (suite *StudentSuite) TestStudentRoute() {
	req, _ := http.NewRequest("GET", "/students/", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	expected := `["Alfredo.Pena","Antonio.Margarito","Arthur.Morgan","John.Marston","Michael.Paul","Tony.Gonzalez","Xin.Zhang"]`

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal(expected, suite.Response.Body.String())
}

// test to verify `/students/{id}` functionality
func (suite *StudentSuite) TestStudentByIdRoute() {
	req, _ := http.NewRequest("GET", "/students/Alfredo.Pena", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	expected := `{"Id":"Alfredo.Pena","AverageScore":"1.00","Results":[{"examId":1,"score":"1.00"},{"examId":2,"score":"1.00"}]}`

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal(expected, suite.Response.Body.String())
}

// test to verify `/students/{id}` functionality
func (suite *StudentSuite) TestStudentByIdRouteTwo() {
	req, _ := http.NewRequest("GET", "/students/Arthur.Morgan", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	expected := `{"Id":"Arthur.Morgan","AverageScore":"0.49","Results":[{"examId":1,"score":"0.77"},{"examId":2,"score":"0.20"}]}`

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal(expected, suite.Response.Body.String())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestStudentSuite(t *testing.T) {
	suite.Run(t, new(StudentSuite))
}
