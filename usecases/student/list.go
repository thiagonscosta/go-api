package student

import (
	"github.com/go-crud/entities"
)

func List() (students []entities.Student, err error) {
	return entities.Students, err
}