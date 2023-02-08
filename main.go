package main 

import (
	// "github.com/go-crud/shared"

	// "strconv"
	// "net/http"
	// "github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	"log"
	"context"
	"github.com/go-crud/api"
	"github.com/go-crud/infra/config"
	"github.com/go-crud/infra/database"
	"github.com/go-crud/infra/database/mongo"
	"github.com/go-crud/infra/database/mongo/repositories"
)

// func healthCheck(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello, it's me",
// 	})
// }

// type Student struct {
// 	ID uuid.UUID `json:"id"`
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }

// var Students = []Student {
// 	Student{ID: shared.GenID(), Name: "Student 1", Age: 8},
// 	Student{ID: shared.GenID(), Name: "Student 2", Age: 8},
// 	Student{ID: shared.GenID(), Name: "Student 3", Age: 8},
// 	Student{ID: shared.GenID(), Name: "Student 4", Age: 8},
// 	Student{ID: shared.GenID(), Name: "Student 5", Age: 8},
// 	Student{ID: shared.GenID(), Name: "Student 6", Age: 8},
// }

// func getStudents(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": Students,
// 	})
// }

// func createStudent(c *gin.Context) {
// 	var student Student

// 	err := c.Bind(&student)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Request failed",
// 		})
// 		return 
// 	}

// 	student.ID = shared.GenID()
// 	Students = append(Students, student)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Successfully created",
// 		"data": student,
// 	})
// }

// func deleteStudent(c *gin.Context) {
// 	idParam := c.Params.ByName("id")

// 	id, err := shared.GetUUIDFromString(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Erro on generating id",
// 		})
// 	}

// 	var students []Student	
	
// 	for _, el := range Students {
// 		if el.ID != id {
// 			students = append(students, el)
// 		} else {
// 			continue 
// 		}
// 	}

// 	Students = students

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Register deleted with success",
// 		"data": Students,
// 	})
// }

// func updateStudent(c *gin.Context) {
// 	var payload Student
// 	var studentLocal Student 
// 	var newStudents []Student

// 	err := c.BindJSON(&payload)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Bad request",
// 		})
// 		return 
// 	}

// 	idParam := c.Params.ByName("id")
// 	id, err := shared.GetUUIDFromString(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error on generating id",
// 		})
// 		return 
// 	}

// 	for _, el := range Students {
// 		if el.ID == id {
// 			studentLocal = el
// 		}
// 	}

// 	if studentLocal.ID == shared.GetUUIDEmpty() {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Student not found",
// 		})
// 		return 
// 	}

// 	studentLocal.Name = payload.Name	
// 	studentLocal.Age = payload.Age
	
// 	for _, el := range Students {
// 		if id == el.ID {
// 			newStudents = append(newStudents, studentLocal)
// 		} else {
// 			newStudents = append(newStudents, el)
// 		}
// 	}

// 	Students = newStudents
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Student updated successfully",
// 		"data": studentLocal,
// 	})
// }

// func getStudent(c *gin.Context) {
// 	idParam := c.Params.ByName("id")
// 	id, err := shared.GetUUIDFromString(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Error on generating id",
// 		})
// 	}
// 	var student Student

// 	for _, el := range Students {
// 		if el.ID == id {
// 			student = el
// 		}
// 	}

// 	if student.ID == shared.GetUUIDEmpty() {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Student not found",
// 		})
// 		return 
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": student,
// 	})
// }

func main() {
	// service := api.NewService()
	// service.Start()
	// service := gin.Default()
	var err error 
	// ctx := context.TODO()

	// db := GetDatabase(ctx)

	err = config.StartConfig()
	FatalError(err)

	err = api.NewService().Start()
	FatalError(err)

	// getRoutes(service)

	// service.Run()
}

// func getRoutes(c *gin.Engine) *gin.Engine {
// 	c.GET("/healthcheck", healthCheck)

// 	// students
// 	groupStudents := c.Group("/students")
// 	groupStudents.GET("/", getStudents)
// 	groupStudents.POST("/", createStudent)
// 	groupStudents.PUT("/:id", updateStudent)
// 	groupStudents.DELETE("/:id", deleteStudent)
// 	groupStudents.GET("/:id", getStudent)

// 	return c
// }

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	studentRepository := repositories.NewStudentRepository(ctx, client) 
	return database.NewDatabase(client, studentRepository)
}