package controllers

// student_exam.go

import (
	database "erik-craigo-code-challenge/db"
	"erik-craigo-code-challenge/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
)

type ExamController struct {
	ExamsDB        database.Exam
	StudentExamsDB database.StudentExam
}

// constructor
func NewExamController(db *memdb.MemDB) ExamController {
	return ExamController{
		ExamsDB:        database.Exam{DB: db},
		StudentExamsDB: database.StudentExam{DB: db},
	}
}

// List Godoc
// @Summary A REST API `/exams` that lists all the exams that have been recorded
// @Param id path int true "Exam ID"
// @Produce  json
// @Router /exams [get]
func (e ExamController) List(c *gin.Context) {
	// get all exams from database
	exams := e.ExamsDB.GetAll()
	// return exams
	c.JSON(http.StatusOK, exams)
}

type GetByNumberOut struct {
	ExamId       int                  `json:"exam_id"`
	AverageScore float64              `json:"average_score"`
	StudentExams []models.StudentExam `json:"student_exams"`
}

// GetByNumber godoc
// @Summary A REST API `/exams/{number}` that lists all the results for the specified exam, and provides the average score across all students
// @Param id path int true "Exam ID"
// @Produce  json
// @Router /exams [get]
func (e ExamController) GetByNumber(c *gin.Context) {
	// get exam number from request
	id, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get studentExams from database
	studentExams := e.StudentExamsDB.GetByExamId(id)

	totalScore := 0.0
	for _, studentExam := range studentExams {
		// string to float
		studentExamScore, _ := strconv.ParseFloat(studentExam.Score, 64)
		totalScore += studentExamScore
	}
	averageScore := float64(totalScore) / float64(len(studentExams))
	// trim average score to 2 decimal places
	averageScore = float64(int(averageScore*100)) / 100

	out := GetByNumberOut{
		ExamId:       id,
		AverageScore: averageScore,
		StudentExams: studentExams,
	}

	// return exam
	c.JSON(http.StatusOK, out)
}

// HandleUpdateExamFromEvent godoc
// @Summary Inserts an exam into the database from an event
func (se ExamController) HandleUpdateExamFromEvent(examId int) {
	// score to string
	// check if Exam in database
	if se.ExamsDB.Exists(examId) {
		// get existing exam data
		return
	} else {

		exam := models.Exam{
			Id: examId,
		}

		// insert into database
		se.ExamsDB.Insert(exam)
	}
}
