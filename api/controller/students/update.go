package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/api/controller"

	"github.com/go-crud/entities"
	"github.com/go-crud/entities/shared"
	student_usecases "github.com/go-crud/usecases/student"
)

func Update(c *gin.Context) {
	var input StudentInput
	var err error

	err = c.Bind(&input)
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

	student, err := student_usecases.Update(input.UUID, input.Name, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	c.JSON(http.StatusOK, studentFound)
}