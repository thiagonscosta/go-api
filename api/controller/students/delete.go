package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/go-crud/entities"
	"github.com/go-crud/entities/shared"
	"github.com/go-crud/api/controller"
	student_usecases "github.com/go-crud/usecases/student"
)

func Delete(c *gin.Context) {
	var input StudentInput

	var err error 

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUUIDFromString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Error on parsing id"))
		return 
	}

	if err = student_usecases.Delete(input.UUID); err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return 
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Student excluding with success"))
}