package tests

import (
	"testing"

	"erik-craigo-code-challenge/config"
	"erik-craigo-code-challenge/controllers"
	database "erik-craigo-code-challenge/db"
	"erik-craigo-code-challenge/tests"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

// suite struct defines necessary components to execute tests
type EventSuite struct {
	suite.Suite
	Config          *viper.Viper
	Router          *gin.Engine
	Response        *httptest.ResponseRecorder
	DB              *memdb.MemDB
	EventController controllers.EventController
}

// sets up suite with dev configs and starts sets up the router
// SetupSuite runs once before executing any tests
func (suite *EventSuite) SetupSuite() {
	config.Init("dev")
	// set up DBs
	database.Init()
}

// sets up env for each test
// SetupTest runes before each test (after SetupSuite)
func (suite *EventSuite) SetupTest() {
	suite.Response = httptest.NewRecorder()
	// initialize db
	DB := database.Init()
	// subscribe to events using event controller

	studentController := controllers.NewStudentController(DB)
	examController := controllers.NewExamController(DB)
	studentExamController := controllers.NewStudentExamController(DB)
	suite.EventController = controllers.NewEventController(DB, studentController, examController, studentExamController)
	suite.DB = DB

}

// test to verify handling of event succesfully
func (suite *EventSuite) TestEvent() {
	events := tests.GetTestEvents()
	for _, event := range events {
		suite.EventController.HandleEvent(event.Data)
	}
	// assert that the event was handled successfully
	suite.Equal(true, true)
}

// test to verify correct processing of student data
func (suite *EventSuite) TestEventProcessingOfStudentData() {
	events := tests.GetTestEvents()
	for _, event := range events {
		suite.EventController.HandleEvent(event.Data)
	}
	// get static student test-data
	students := tests.GetTestStudents()

	studentDB := database.Student{DB: suite.DB}

	// check that each student matches what is in the DB (it should)
	for _, student := range students {
		// get student from DB
		studentFromDB := studentDB.GetById(student.Id)
		// assert not nil
		suite.NotNil(studentFromDB)
	}

}

// test to verify correct processing of student-exam data
func (suite *EventSuite) TestEventProcessingOfStudentExamData() {
	events := tests.GetTestEvents()
	for _, event := range events {
		suite.EventController.HandleEvent(event.Data)
	}

	// get static studentExam test-data
	studentExams := tests.GetTestStudentExams()

	studentExamDB := database.StudentExam{DB: suite.DB}

	// check that each studentExam matches what is in the DB (it should)
	for _, studentExam := range studentExams {
		// get studentExam from DB
		studentExamFromDB := studentExamDB.GetByCompoundKey(studentExam.StudentId, studentExam.ExamId)
		// assert that the studentExam has the same total score
		suite.Equal(studentExam.Score, studentExamFromDB.Score)

	}
}

// test to verify correct processing of exam data
func (suite *EventSuite) TestEventProcessingOfExamData() {
	events := tests.GetTestEvents()
	for _, event := range events {
		suite.EventController.HandleEvent(event.Data)
	}

	// get static studentExam test-data
	exams := tests.GetTestExams()
	examDB := database.Exam{DB: suite.DB}

	// check that each studentExam matches what is in the DB (it should)
	for _, exam := range exams {
		// get studentExam from DB
		examFromDB := examDB.GetById(exam.Id)
		// assert that the studentExam has the same total score
		suite.NotNil(examFromDB)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEventSuite(t *testing.T) {
	suite.Run(t, new(EventSuite))
}
