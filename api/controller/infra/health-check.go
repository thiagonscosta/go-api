package infra

import (
	"net/http"
	"github.com/gin-gonic/gin"
	response "github.com/go-crud/api/controller"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewResponseMessage("Ok"))
}