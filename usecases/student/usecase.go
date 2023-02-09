package student

import "github.com/go-crud/infra/database"

type StudentUseCase struct {
	Database *database.Database
}

func NewStudentUsecase(db *database.Database) *StudentUseCase {
	return &StudentUseCase{
		Database: db,
	}
}