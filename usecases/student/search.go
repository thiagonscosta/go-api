package student

import (
	"errors"
	"github.com/google/uuid"
	"github.com/go-crud/entities"

	"github.com/go-crud/entities/shared"
)

func SearchById(id uuid.UUID) (student entities.Student, err error) {
	var studentFound entities.Student

	for _, el := range entities.Students {
		if el.ID == id {
			studentFound = el
		}
	}

	if studentFound.ID == shared.GetUUIDEmpty() {
		return studentFound, errors.New("Student not found")
	}

	return studentFound, err 
}