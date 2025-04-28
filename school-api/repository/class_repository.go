package repository

import (
	"school-api/models"
	"gorm.io/gorm"
)

type ClassRepository interface {
	Create(class *models.Class) error
	GetAll() ([]models.Class, error)
	GetByID(id uint) (*models.Class, error)
	Update(class *models.Class) error
	Delete(id uint) error
}

type classRepository struct {
	GenericRepository[models.Class]
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{
		GenericRepository: NewGenericRepository[models.Class](db),
	}
} 