package service

import (
	"school-api/models"
	"school-api/repository"
)

type StudentService interface {
	CreateStudent(student *models.Student) error
	GetAllStudents() ([]models.Student, error)
	GetStudentByID(id uint) (*models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{
		studentRepo: studentRepo,
	}
}

func (s *studentService) CreateStudent(student *models.Student) error {
	return s.studentRepo.Create(student)
}

func (s *studentService) GetAllStudents() ([]models.Student, error) {
	return s.studentRepo.GetAll()
}

func (s *studentService) GetStudentByID(id uint) (*models.Student, error) {
	return s.studentRepo.GetByID(id)
}

func (s *studentService) UpdateStudent(student *models.Student) error {
	return s.studentRepo.Update(student)
}

func (s *studentService) DeleteStudent(id uint) error {
	return s.studentRepo.Delete(id)
} 