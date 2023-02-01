package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/api/controller"

	"github.com/go-crud/entities"
	"github.com/go-crud/entities/shared"
)

func Update(c *gin.Context) {
	var input StudentInput
	var studentFound entities.Student
	var newStudents []entities.Student

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUUIDFromString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	for _, el := range entities.Students {
		if el.ID == input.UUID {
			studentFound = el
		}
	}

	if studentFound.ID == shared.GetUUIDEmpty() {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return 
	}

	studentFound.Name = input.Name
	studentFound.Age = input.Age
	
	for _, el := range entities.Students {
		if studentFound.ID == el.ID {
			newStudents = append(newStudents, studentFound)
		} else {
			newStudents = append(newStudents, el)
		}
	}

	entities.Students = newStudents

	c.JSON(http.StatusOK, studentFound)
}