package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-crud/entities"
)

type Database struct {
	Conn *mongo.Client

	StudentRepository entities.StudentRepository
}

func NewDatabase(conn *mongo.Client, str entities.StudentRepository) *Database {
	return &Database{
		Conn: conn,
		StudentRepository: str,
	}
}