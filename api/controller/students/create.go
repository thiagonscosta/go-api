package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/entities"

	"github.com/go-crud/api/controller"
)

func Create (c *gin.Context) {
	var input StudentInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
	}

	student := entities.NewStudent(input.Name, input.Age)
	entities.Students = append(entities.Students, *student)

	c.JSON(http.StatusAccepted, student)
}