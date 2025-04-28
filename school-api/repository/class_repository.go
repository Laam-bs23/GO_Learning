package repository

import (
	"school-api/models"
	"gorm.io/gorm"
)

type ClassRepository interface {
	Create(class *models.Class) error
	GetAll() ([]models.Class, error)
	GetByID(id uint) (*models.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Create(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *classRepository) GetAll() ([]models.Class, error) {
	var classes []models.Class
	err := r.db.Find(&classes).Error
	return classes, err
}

func (r *classRepository) GetByID(id uint) (*models.Class, error) {
	var class models.Class
	err := r.db.First(&class, id).Error
	if err != nil {
		return nil, err
	}
	return &class, nil
} 