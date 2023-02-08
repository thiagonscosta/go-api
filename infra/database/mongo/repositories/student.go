package repositories

import (
	"context"
	mongo_drive "go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	Ctx context.Context
	Collection mongo_drive.Collection
}

func NewStudentRepository(ctx context.Context) *StudentRepository {
	return &StudentRepository {
		Ctx: ctx,
	}
}