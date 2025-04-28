package service

import (
	"school-api/models"
	"school-api/repository"
)

type ClassService interface {
	CreateClass(class *models.Class) error
	GetAllClasses() ([]models.Class, error)
	GetClassByID(id uint) (*models.Class, error)
}

type classService struct {
	repo repository.ClassRepository
}

func NewClassService(repo repository.ClassRepository) ClassService {
	return &classService{repo: repo}
}

func (s *classService) CreateClass(class *models.Class) error {
	return s.repo.Create(class)
}

func (s *classService) GetAllClasses() ([]models.Class, error) {
	return s.repo.GetAll()
}

func (s *classService) GetClassByID(id uint) (*models.Class, error) {
	return s.repo.GetByID(id)
} 