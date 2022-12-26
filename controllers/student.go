package controllers

// student.go

import (
	database "erik-craigo-code-challenge/db"
	"erik-craigo-code-challenge/models"
	"net/http"
	"strconv"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
)

// StudentController godoc

type StudentController struct {
	StudentsDB     database.Student
	StudentExamsDB database.StudentExam
}

// constructor
func NewStudentController(db *memdb.MemDB) StudentController {
	return StudentController{
		StudentsDB:     database.Student{DB: db},
		StudentExamsDB: database.StudentExam{DB: db},
	}
}

// eventToStudent godoc
// @Summary produces a student object from an event object
// @Description produces a student object from an event object
// @param event models.Event object
// @produces student models.Student object
func eventToStudent(event models.RawData) models.Student {
	// create student object
	student := models.Student{
		Id: event.StudentId,
	}

	// return student object
	return student
}

// Upsert godoc
// Optionally insert user (student) into database
// @Param c gin.Context object
// @Param event models.Event object
// @Summary Responds with boolean value if user is not already in database
// @Description Optionally inserts user into database
// @Produce  json
func (s StudentController) Upsert(c *gin.Context, event models.RawData) bool {
	// check if student is in database already, return false if so
	if s.Exists(event.StudentId) {
		return false
	}
	// create student object
	student := eventToStudent(event)
	// insert student into database
	// use student model to insert student into database
	// return true if student is inserted into database
	return s.StudentsDB.Insert(student)

}

// Exists godoc
// @Param c gin.Context object
// @Param studentID string
// @Summary Responds with boolean value if user is not already in database
// @Description Optionally inserts user into database
// @Produce  json
func (s StudentController) Exists(studentID string) bool {
	// check if student is in database already, return false if so
	// return true if student is in database
	return s.StudentsDB.Exists(studentID)
}

// Returns a list of students in the database
// @Param c gin.Context object
// @Summary Responds with boolean value if user is not already in database
// @Description Optionally inserts user into database
// @Produce  json
// @Router /students [get]
func (s StudentController) Get(c *gin.Context) {
	// get students from database
	// use student model to get students from database
	// return students
	students := s.StudentsDB.GetAll()
	// create list of just student ids
	studentIds := []string{}
	for _, student := range students {
		studentIds = append(studentIds, student.Id)
	}
	c.JSON(http.StatusOK, studentIds)
}

type Result struct {
	ExamId int    `json:"examId"`
	Score  string `json:"score"`
}

type GetByIdOut struct {
	Id           string   `json id`
	AverageScore string   `json averageScore`
	Results      []Result `json results`
}

// Returns a list of students in the database
// @Param c gin.Context object
// @Param studentID string
// @Summary Responds with boolean value if user is not already in database
// @Description Optionally inserts user into database
// @Success 200 {string} Working!
// @Failure 404 {string} Not Found!
// @Produce  json
func (s StudentController) GetById(c *gin.Context) {
	// get students from database
	// use student model to get students from database
	// return students
	studentId := c.Param("id")
	student := s.StudentsDB.GetById(studentId)
	studentExams := s.StudentExamsDB.GetByStudentId(studentId)
	results := []Result{}
	totalScore := 0.0
	count := 0.0
	println(len(studentExams))
	for _, studentExam := range studentExams {
		// score to float
		score, _ := strconv.ParseFloat(studentExam.Score, 64)

		result := Result{
			ExamId: studentExam.ExamId,
			Score:  fmt.Sprintf("%.2f", score),
		}
		results = append(results, result)

		totalScore += score
		count += 1
	}
	averageScore := fmt.Sprintf("%.2f", totalScore/count)

	jsonToEmit := GetByIdOut{
		Id:           student.Id,
		AverageScore: averageScore,
		Results:      results,
	}
	// marshal json
	c.JSON(http.StatusOK, jsonToEmit)
}

// GetStudentExams godoc
// @Summary lists all the exams that the specified student has taken, and provides the student's average score across all exams
// @Param id path int true "Student ID"
// @Produce  json
// @Router /students/{studentID}/exams [get]
func (s StudentController) GetStudentExams(c *gin.Context, studentID string) {
	// get student exams from database
	// use student model to get student exams from database
	// return student exams
	studentExams := s.StudentExamsDB.GetByStudentId(studentID)
	c.JSON(http.StatusOK, studentExams)
}

// Function which adds a user to the database, updates the users score
// @Summary adds a user to the database, updates the user's score
// @Param studentID string
// @Param parScore string
func (s StudentController) HandleUpdateUserFromEvent(studentId string) {

	// check if student is in database already
	if s.Exists(studentId) {
		// insert student into database
		// use student model to insert student into database
		return
	} else {

		student := models.Student{
			Id: studentId,
		}
		s.StudentsDB.Insert(student)
	}
}
