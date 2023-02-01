package student 

import (
	"github.com/google/uuid"
	"github.com/go-crud/entities"
)

func Delete(id uuid.UUID) (err error){
	var newStudents []entities.Student

	for _, el := range entities.Students {
		if el.ID != id {
			newStudents = append(newStudents, el) 
		}
	}

	entities.Students = newStudents
}