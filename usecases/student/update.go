package student 

import (
	"errors"
	"github.com/google/uuid"

	"github.com/go-crud/entities"
	"github.com/go-crud/entities/shared"
)

func Update(id uuid.UUID, name string, age int) (student entities.Student, err error) {
	var studentFound entities.Student
	var newStudents []entities.Student

	for _, el := range entities.Students {
		if el.ID == id {
			studentFound = el
		}
	}

	if studentFound.ID == shared.GetUUIDEmpty() {
		return student, errors.New("Student not found")
	}

	studentFound.Name = name
	studentFound.Age = age

	for _, el := range entities.Students {
		if studentFound.ID == el.ID {
			newStudents = append(newStudents, studentFound)
		} else {
			newStudents = append(newStudents, el)
		}
	}

	entities.Students = newStudents

	return student, err
}