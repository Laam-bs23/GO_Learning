package repository

import (
	"school-api/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *models.Student) error
	GetAll() ([]models.Student, error)
	GetByID(id uint) (*models.Student, error)
	Update(student *models.Student) error
	Delete(id uint) error
}

type studentRepository struct {
	GenericRepository[models.Student]
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		GenericRepository: NewGenericRepository[models.Student](db),
	}
} 