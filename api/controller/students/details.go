package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/entities"
	"github.com/go-crud/api/controller"

	"github.com/google/uuid"

	"github.com/go-crud/entities/shared"
	student_usecases "github.com/go-crud/usecases/student"
)

func Details(c *gin.Context) {
	var input StudentInput 
	var student entities.Student
	var err error

	input.ID = c.Params.ByName("id")	
	input.UUID, err = shared.GetUUIDFromString(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Error on parsing parameters"))
		return
	}

	student, err = student_usecases.SearchById(input.UUID)

	if student.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Not found"))
		return
	}

	c.JSON(http.StatusOK, student)
}