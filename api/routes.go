package api

import (
	infra_controller "github.com/go-crud/api/controller/infra"
	students_controller "github.com/go-crud/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/healthcheck", infra_controller.HealthCheck)

	groupStudents := s.Engine.Group("students")
	groupStudents.GET("/", students_controller.List)
	groupStudents.POST("/", students_controller.Create)
	groupStudents.PUT("/:id", students_controller.Update)
	groupStudents.DELETE("/:id", students_controller.Delete)
}