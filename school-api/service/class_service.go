package service

import (
	"school-api/models"
	"school-api/repository"
)

type ClassService interface {
	CreateClass(class *models.Class) error
	GetAllClasses() ([]models.Class, error)
	GetClassByID(id uint) (*models.Class, error)
	UpdateClass(class *models.Class) error
	DeleteClass(id uint) error
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

func (s *classService) UpdateClass(class *models.Class) error {
	return s.repo.Update(class)
}

func (s *classService) DeleteClass(id uint) error {
	return s.repo.Delete(id)
} 