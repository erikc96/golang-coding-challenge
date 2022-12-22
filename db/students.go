package database

import (
	// uuid
	"erik-craigo-code-challenge/models"

	"github.com/hashicorp/go-memdb"
)

type Student struct {
	DB *memdb.MemDB
}

// Students

// add student
// Insert is used to add or update an object into the given table.
func (s Student) Insert(student models.Student) bool {
	db := s.DB
	txn := db.Txn(true)
	defer txn.Commit()

	if err := txn.Insert("students", student); err != nil {
		panic(err)
		return false
	}
	return true
}

// delete student
func (s Student) Delete(id string) {
	db := s.DB
	txn := db.Txn(true)
	defer txn.Commit()

	if err := txn.Delete("students", id); err != nil {
		panic(err)
	}
}

// get all students
func (s Student) GetAll() []models.Student {
	db := s.DB
	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("students", "id")
	if err != nil {
		panic(err)
	}

	var students []models.Student
	for {
		raw := it.Next()
		if raw == nil {
			break
		}
		student := raw.(models.Student)
		students = append(students, student)
	}
	return students
}

// get student by id
func (s Student) GetById(id string) models.Student {
	db := s.DB
	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.First("students", "id", id)
	if err != nil {
		panic(err)
	}

	return it.(models.Student)
}

// Exists godoc
// @Param studentID string
// @Summary Responds with boolean value
// @Description Optionally inserts user into database
// @Produce  boolean
func (s Student) Exists(studentID string) bool {
	db := s.DB
	result, err := db.Txn(false).First("students", "id", studentID)
	if err != nil {
		panic(err)
	}
	return result != nil
}
