package students 

import (
	student_usecase "github.com/go-crud/usecases/student"
)

type StudentController struct {
	StudentUseCase *student_usecase.StudentUseCase
}

func NewStudentController(su *student_usecase.StudentUseCase) *StudentController {
	return &StudentController {
		StudentUseCase: su,
	}
}