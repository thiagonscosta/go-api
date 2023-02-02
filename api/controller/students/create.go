package students

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-crud/api/controller"
	student_usecases "github.com/go-crud/usecases/student"
)

func Create (c *gin.Context) {
	var input StudentInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
	}

	// student := entities.NewStudent(input.Name, input.Age)
	// entities.Students = append(entities.Students, *student)
	student, err := student_usecases.Create(input.Name, input.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
	}
	c.JSON(http.StatusAccepted, student)
}