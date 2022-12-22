package server

import (
	"erik-craigo-code-challenge/controllers"
	database "erik-craigo-code-challenge/db"
	"log"
)

func Init() {
	// initialize db
	DB := database.Init()
	// subscribe to events using event controller

	studentController := controllers.NewStudentController(DB)
	examController := controllers.NewExamController(DB)
	studentExamController := controllers.NewStudentExamController(DB)
	eventController := controllers.NewEventController(DB, studentController, examController, studentExamController)

	r := SetupRouter(studentController, studentExamController, examController)
	// run event controller in goroutine
	go eventController.Listen()

	err := r.Run(":3000")

	if err != nil {
		log.Fatal(err)
	}
}
