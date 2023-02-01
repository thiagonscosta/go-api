package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/entities"
	"github.com/go-crud/entities/shared"
	"github.com/go-crud/api/controller"
)

func Delete(c *gin.Context) {
	idParam := c.Params.ByName("id")
	var students []entities.Student

	id, err := shared.GetUUIDFromString(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return 
	}

	for _, el := range entities.Students {
		if el.ID != id {
			students = append(students, el)
		} else {
			continue
		}
	}

	entities.Students = students

	c.JSON(http.StatusOK, students)
}