package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/entities"
	"github.com/go-crud/api/controller"

	"github.com/go-crud/entities/shared"
)

func Details(c *gin.Context) {
	var input StudentInput 
	var student entities.Student
	var err error

	input.ID = c.Params.ByName("id")	
	input.UUID, err = shared.GetUUIDFromString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Error no parsing parameters"))
		return
	}

	for _, el := range entities.Students {
		if el.ID == input.UUID {
			student = el
		}
	}

	// if student.ID == shared.GetUUIDEmpty() {
	// 	c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Not found"))
	// 	return
	// }

	c.JSON(http.StatusOK, student)
}