package database

import (
	// uuid
	"erik-craigo-code-challenge/models"

	"github.com/hashicorp/go-memdb"
)

// Student Exams

type StudentExam struct {
	DB *memdb.MemDB
}

// insert student-exam
// Insert is used to add or update an object into the given table.
func (se StudentExam) Insert(studentExam models.StudentExam) {
	db := se.DB
	txn := db.Txn(true)
	defer txn.Commit()
	if err := txn.Insert("studentExams", studentExam); err != nil {
		panic(err)
	}
}

// get by compount key
func (se StudentExam) GetByCompoundKey(StudentId string, ExamId int) models.StudentExam {
	db := se.DB
	txn := db.Txn(false)
	defer txn.Abort()

	studentExam, err := txn.First("studentExams", "key", StudentId, ExamId)

	if err != nil {
		panic(err)
	}

	return studentExam.(models.StudentExam)
}

// get student-exam by studentId
func (se StudentExam) GetByStudentId(studentId string) []models.StudentExam {
	db := se.DB
	txn := db.Txn(false)
	defer txn.Abort()

	iterator, err := txn.Get("studentExams", "studentId", studentId)
	if err != nil {
		panic(err)
	}

	var studentExams []models.StudentExam
	for {
		raw := iterator.Next()
		if raw == nil {
			break
		}
		studentExam := raw.(models.StudentExam)
		studentExams = append(studentExams, studentExam)
	}

	return studentExams
}

// get student-exam by examId
func (se StudentExam) GetByExamId(examId int) []models.StudentExam {
	db := se.DB
	txn := db.Txn(false)
	defer txn.Abort()

	iterator, err := txn.Get("studentExams", "examId", examId)
	if err != nil {
		panic(err)
	}

	var studentExams []models.StudentExam
	for {
		raw := iterator.Next()
		if raw == nil {
			break
		}
		studentExams = append(studentExams, raw.(models.StudentExam))
	}

	return studentExams
}

// get all
func (se StudentExam) GetAll() []models.StudentExam {
	db := se.DB
	txn := db.Txn(false)
	defer txn.Abort()

	iterator, err := txn.Get("studentExams", "id")
	if err != nil {
		panic(err)
	}

	var studentExams []models.StudentExam
	for {
		raw := iterator.Next()
		if raw == nil {
			break
		}
		studentExams = append(studentExams, raw.(models.StudentExam))
	}

	return studentExams
}
