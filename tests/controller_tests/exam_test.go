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
type ExamSuite struct {
	suite.Suite
	Config   *viper.Viper
	Router   *gin.Engine
	Response *httptest.ResponseRecorder
}

// sets up suite with dev configs and starts sets up the router
// SetupSuite runs once before executing any tests
func (suite *ExamSuite) SetupSuite() {
	config.Init("dev")
	suite.Config = config.GetConfig()
	// set up DBs
	DB := database.Init()

	studentController := controllers.NewStudentController(DB)
	examController := controllers.NewExamController(DB)
	studentExamController := controllers.NewStudentExamController(DB)

	suite.Router = server.SetupRouter(studentController, studentExamController, examController)
	examDB := database.Exam{DB}
	studentExamDB := database.StudentExam{DB}

	// get student test data
	exams := tests.GetTestExams()
	studentExams := tests.GetTestStudentExams()
	// populate DBs
	for _, exam := range exams {
		examDB.Insert(exam)
	}
	for _, studentExam := range studentExams {
		studentExamDB.Insert(studentExam)
	}
}

// sets up env for each test
// SetupTest runes before each test (after SetupSuite)
func (suite *ExamSuite) SetupTest() {
	suite.Response = httptest.NewRecorder()
}

// test to verify `/exams/` functionality
func (suite *ExamSuite) TestExamsRoute() {
	req, _ := http.NewRequest("GET", "/exams/", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	expected := `[{"id":1},{"id":2},{"id":3},{"id":256}]`

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal(expected, suite.Response.Body.String())
}

// test to verify `/exams/:number` functionality
func (suite *ExamSuite) TestExamsByNumberRoute() {
	req, _ := http.NewRequest("GET", "/exams/1", nil)
	suite.Router.ServeHTTP(suite.Response, req)

	expected := `{"exam_id":1,"average_score":0.91,"student_exams":[{"id":6,"StudentId":"Alfredo.Pena","ExamId":1,"Score":"1.000000"},{"id":2,"StudentId":"Arthur.Morgan","ExamId":1,"Score":"0.770379"},{"id":1,"StudentId":"John.Marston","ExamId":1,"Score":"0.900000"},{"id":0,"StudentId":"Michael.Paul","ExamId":1,"Score":"1.000000"}]}`

	suite.Equal(http.StatusOK, suite.Response.Code)
	suite.Equal(expected, suite.Response.Body.String())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExamSuite(t *testing.T) {
	suite.Run(t, new(ExamSuite))
}
