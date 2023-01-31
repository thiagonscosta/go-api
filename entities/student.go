package entities

import (
	"github.com/google/uuid"
	"github.com/go-crud/entities/shared"
)

type Student struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var Students = []Student {
	Student{ID: shared.GenID(), Name: "Student 1", Age: 8},
	Student{ID: shared.GenID(), Name: "Student 2", Age: 8},
	Student{ID: shared.GenID(), Name: "Student 3", Age: 8},
	Student{ID: shared.GenID(), Name: "Student 4", Age: 8},
	Student{ID: shared.GenID(), Name: "Student 5", Age: 8},
	Student{ID: shared.GenID(), Name: "Student 6", Age: 8},
}

func NewStudent(name string, age int) *Student {
	return &Student{
		ID: shared.GenID(),
		Name: name,
		Age: age,
	}
}
