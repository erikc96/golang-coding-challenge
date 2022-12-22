package database

import (
	// uuid
	"sync"

	"github.com/hashicorp/go-memdb"
)

var dbInstance *memdb.MemDB

var lock = &sync.Mutex{}

func Init() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"students": {
				Name: "students",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
				},
			},
			"exams": {
				Name: "exams",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
				},
			},
			"studentExams": {
				Name: "studentExams",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:   "id",
						Unique: true,
						Indexer: &memdb.CompoundIndex{
							Indexes: []memdb.Indexer{
								&memdb.StringFieldIndex{Field: "StudentId"},
								&memdb.IntFieldIndex{Field: "ExamId"},
								&memdb.StringFieldIndex{Field: "Score"},
							},
						},
					},
					"key": {
						Name:   "key",
						Unique: false,
						Indexer: &memdb.CompoundIndex{
							Indexes: []memdb.Indexer{
								&memdb.StringFieldIndex{Field: "StudentId"},
								&memdb.IntFieldIndex{Field: "ExamId"},
							},
						},
					},
					"studentId": {
						Name:    "studentId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "StudentId"},
					},
					"examId": {
						Name:    "examId",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "ExamId"},
					},
					"score": {
						Name:    "score",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Score"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Commit the transaction
	txn.Commit()
	dbInstance = db
	return db
}

// Define the DB schema

func GetDbInstance() *memdb.MemDB {
	return dbInstance
}
