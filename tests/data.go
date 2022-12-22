package tests

import "erik-craigo-code-challenge/models"

var (
	student1 = models.Student{
		Id: "Michael.Paul",
	}
	student2 = models.Student{
		Id: "John.Marston",
	}
	student3 = models.Student{
		Id: "Arthur.Morgan",
	}
	student4 = models.Student{
		Id: "Alfredo.Pena",
	}
	student5 = models.Student{
		Id: "Antonio.Margarito",
	}
	student6 = models.Student{
		Id: "Tony.Gonzalez",
	}
	student7 = models.Student{
		Id: "Xin.Zhang",
	}
	studentExam1 = models.StudentExam{
		StudentId: "Michael.Paul",
		Id:        0,
		ExamId:    1,
		Score:     "1.000000",
	}
	studentExam2 = models.StudentExam{
		StudentId: "John.Marston",
		Id:        1,
		ExamId:    1,
		Score:     "0.900000",
	}
	studentExam3 = models.StudentExam{
		StudentId: "Arthur.Morgan",
		Id:        2,
		ExamId:    1,
		Score:     "0.770379",
	}
	studentExam4 = models.StudentExam{
		StudentId: "Michael.Paul",
		Id:        3,
		ExamId:    256,
		Score:     "0.800000",
	}
	studentExam5 = models.StudentExam{
		StudentId: "John.Marston",
		Id:        4,
		ExamId:    2,
		Score:     "0.662064",
	}
	studentExam6 = models.StudentExam{
		StudentId: "Arthur.Morgan",
		Id:        5,
		ExamId:    2,
		Score:     "0.200000",
	}
	studentExam7 = models.StudentExam{
		StudentId: "Alfredo.Pena",
		Id:        6,
		ExamId:    1,
		Score:     "1.000000",
	}
	studentExam8 = models.StudentExam{
		StudentId: "Alfredo.Pena",
		Id:        7,
		ExamId:    2,
		Score:     "1.000000",
	}
	studentExam9 = models.StudentExam{
		StudentId: "Antonio.Margarito",
		Id:        8,
		ExamId:    3,
		Score:     "0.505323",
	}
	studentExam10 = models.StudentExam{
		StudentId: "Tony.Gonzalez",
		Id:        9,
		ExamId:    3,
		Score:     "0.000000",
	}
	studentExam11 = models.StudentExam{
		StudentId: "Xin.Zhang",
		ExamId:    3,
		Id:        10,
		Score:     "0.354290",
	}
	studentExam12 = models.StudentExam{
		StudentId: "Xin.Zhang",
		ExamId:    2,
		Id:        11,
		Score:     "0.423327",
	}
	studentExam13 = models.StudentExam{
		StudentId: "Tony.Gonzalez",
		ExamId:    2,
		Id:        12,
		Score:     "0.345253",
	}
	exam1 = models.Exam{
		Id: 1,
	}
	exam2 = models.Exam{
		Id: 2,
	}
	exam3 = models.Exam{
		Id: 3,
	}
	exam256 = models.Exam{
		Id: 256,
	}

	students = []models.Student{
		student1,
		student2,
		student3,
		student4,
		student5,
		student6,
		student7,
	}

	studentExams = []models.StudentExam{
		studentExam1,
		studentExam2,
		studentExam3,
		studentExam4,
		studentExam5,
		studentExam6,
		studentExam7,
		studentExam8,
		studentExam9,
		studentExam10,
		studentExam11,
		studentExam12,
		studentExam13,
	}
)
var (
	Event1 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Michael.Paul",
			ExamId:    1,
			Score:     1.00,
		},
	}
	event2 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "John.Marston",
			ExamId:    1,
			Score:     0.90,
		},
	}
	event3 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Arthur.Morgan",
			ExamId:    1,
			Score:     0.7703790745194544,
		},
	}
	event4 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Michael.Paul",
			ExamId:    256,
			Score:     0.80,
		},
	}
	event5 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "John.Marston",
			ExamId:    2,
			Score:     0.6620636142806398,
		},
	}
	event6 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Arthur.Morgan",
			ExamId:    2,
			Score:     0.20,
		},
	}
	event7 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Alfredo.Pena",
			ExamId:    1,
			Score:     1.0,
		},
	}
	event8 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Alfredo.Pena",
			ExamId:    2,
			Score:     1.00,
		},
	}
	event9 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Antonio.Margarito",
			ExamId:    3,
			Score:     0.5053225,
		},
	}
	event10 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Tony.Gonzalez",
			ExamId:    3,
			Score:     0,
		},
	}
	event11 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Xin.Zhang",
			ExamId:    3,
			Score:     0.354290253,
		},
	}
	event12 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Xin.Zhang",
			ExamId:    2,
			Score:     .4233271,
		},
	}
	event13 = models.RawEventData{
		Event: "score",
		Data: models.RawData{
			StudentId: "Tony.Gonzalez",
			ExamId:    2,
			Score:     .34525320,
		},
	}
)

func GetTestEvents() []models.RawEventData {
	return []models.RawEventData{
		Event1,
		event2,
		event3,
		event4,
		event5,
		event6,
		event7,
		event8,
		event9,
		event10,
		event11,
		event12,
		event13,
	}
}

func GetTestStudents() []models.Student {
	return []models.Student{
		student1,
		student2,
		student3,
		student4,
		student5,
		student6,
		student7,
	}
}

func GetTestStudentExams() []models.StudentExam {
	return []models.StudentExam{
		studentExam1,
		studentExam2,
		studentExam3,
		studentExam4,
		studentExam5,
		studentExam6,
		studentExam7,
		studentExam8,
		studentExam9,
		studentExam10,
		studentExam11,
		studentExam12,
		studentExam13,
	}
}

func GetTestExams() []models.Exam {
	return []models.Exam{
		exam1,
		exam2,
		exam3,
		exam256,
	}
}
