package students

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/api/controller"

	student_usecases "github.com/go-crud/usecases/student"
)

func List(c *gin.Context) {
	students, err := student_usecases.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}
	c.JSON(http.StatusOK, students)
}