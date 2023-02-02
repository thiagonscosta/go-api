package student

import (
	"github.com/go-crud/entities"
)

func Create(name string, age int) (student entities.Student, err error) {
	postStudent := entities.NewStudent(name, age)
	student = *postStudent
	entities.Students = append(entities.Students, student)

	return student, err
}