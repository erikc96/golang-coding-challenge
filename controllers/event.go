package controllers

// handles updating state based on events from the eventsource api
// http://live-test-scores.herokuapp.com/scores

import (
	"encoding/json"
	// import eventsource
	database "erik-craigo-code-challenge/db"
	"erik-craigo-code-challenge/models"

	"github.com/donovanhide/eventsource"
	"github.com/hashicorp/go-memdb"
)

// event source url constant
const eventSourceURL = "http://live-test-scores.herokuapp.com/scores"

// event controller
type EventController struct {
	studentsDB            database.Student
	studentExamsDB        database.StudentExam
	studentController     StudentController
	examController        ExamController
	studentExamController StudentExamController
}

// constructor
func NewEventController(DB *memdb.MemDB, studentController StudentController, examController ExamController, studentExamController StudentExamController) EventController {
	studentsDB := database.Student{DB}
	studentExamsDB := database.StudentExam{DB}

	// create event controller
	eventController := EventController{
		studentsDB:            studentsDB,
		studentExamsDB:        studentExamsDB,
		studentController:     studentController,
		examController:        examController,
		studentExamController: studentExamController,
	}

	return eventController

}

func (e EventController) HandleEvent(event models.RawData) {
	// update student
	e.studentController.HandleUpdateUserFromEvent(event.StudentId)
	// update exam
	e.examController.HandleUpdateExamFromEvent(event.ExamId)
	// update student exam
	e.studentExamController.HandleUpdateStudentExamFromEvent(event.StudentId, event.ExamId, event.Score)
}

// create event listener
// event listener controller
func (e EventController) Listen() {
	// create event listener
	listener, err := eventsource.Subscribe(eventSourceURL, "")

	// listen for events
	//handle err
	if err != nil {
		panic(err)
	}
	for {
		// get event
		eventRaw := <-listener.Events
		// converted nested event object (represented as json string) to event object
		nestedEvent := eventRaw.Data()

		// convert json string to event object
		var event models.RawData
		json.Unmarshal([]byte(nestedEvent), &event)
		// get string representation of event using marshall
		eventString, _ := json.Marshal(event)
		println("Processing event: ", string(eventString))

		e.HandleEvent(event)
	}
}
