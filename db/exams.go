package database

import (
	// uuid
	"erik-craigo-code-challenge/models"

	"github.com/hashicorp/go-memdb"
)

// TODO Exams
type Exam struct {
	DB *memdb.MemDB
}

// Get All
func (e Exam) GetAll() []models.Exam {
	db := e.DB
	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("exams", "id")
	if err != nil {
		panic(err)
	}

	var exams []models.Exam
	for {
		raw := it.Next()

		if raw == nil {
			break
		}

		exam := raw.(models.Exam)
		exams = append(exams, exam)
	}
	return exams
}

// Exams Exists
func (e Exam) Exists(examID int) bool {
	db := e.DB
	result, err := db.Txn(false).First("exams", "id", examID)
	if err != nil {
		panic(err)
	}
	return result != nil
}

// Exams Insert
// Insert is used to add or update an object into the given table.
func (e Exam) Insert(exam models.Exam) {
	db := e.DB
	txn := db.Txn(true)
	defer txn.Commit()

	if err := txn.Insert("exams", exam); err != nil {
		panic(err)
	}
}

// Exams Delete
func (e Exam) Delete(examId int) bool {
	db := e.DB
	txn := db.Txn(true)
	defer txn.Commit()

	if err := txn.Delete("exams", examId); err != nil {
		panic(err)
		return false
	}
	return true
}

// Get By Id
func (e Exam) GetById(examId int) models.Exam {
	db := e.DB
	txn := db.Txn(false)
	defer txn.Abort()

	result, err := txn.First("exams", "id", examId)
	if err != nil {
		panic(err)
	}
	return result.(models.Exam)
}
