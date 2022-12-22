package controllers

// student_exam.go

import (
	database "erik-craigo-code-challenge/db"
	"erik-craigo-code-challenge/models"
	"strconv"

	"github.com/hashicorp/go-memdb"
)

type StudentExamController struct {
	StudentExamsDB database.StudentExam
}

// constructor
func NewStudentExamController(db *memdb.MemDB) StudentExamController {
	return StudentExamController{
		StudentExamsDB: database.StudentExam{DB: db},
	}
}
func (se StudentExamController) HandleUpdateStudentExamFromEvent(studentId string, examId int, parScore float64) {
	// score to string
	score := strconv.FormatFloat(parScore, 'f', 6, 64)
	studentExam := models.StudentExam{
		StudentId: studentId,
		ExamId:    examId,
		Score:     score,
	}
	// insert into database
	se.StudentExamsDB.Insert(studentExam)
}
