package api 

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/go-crud/infra/config"
	// "github.com/go-crud/infra/database"
)

type Service struct {
	Engine *gin.Engine

	Database *database.Database

	StudentsController *students.StudentController
}

func NewService() *Service {
	return &Service{
		Engine: gin.Default(),
		Database: db,
	}
}

func (s *Service) Start() error {
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}