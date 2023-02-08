package repositories

import (
	"context"
	mongo_infra "github.com/go-crud/infra/database/mongo"

	mongo_drive "go.mongodb.org/mongo-driver/mongo"
)

const (
	StudentCollection = "students"
)

type StudentRepository struct {
	Ctx context.Context
	Collection *mongo_drive.Collection
}

func NewStudentRepository(ctx context.Context, client *mongo_drive.Client) *StudentRepository {
	return &StudentRepository {
		Ctx: ctx,
		Collection: mongo_infra.GetCollection(ctx, client, StudentCollection),
	}
}