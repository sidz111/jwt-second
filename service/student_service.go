package service

import (
	"context"
	"errors"

	"github.com/sidz111/jtw-second/models"
	"github.com/sidz111/jtw-second/repository"
	"golang.org/x/crypto/bcrypt"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	GetStudentByID(ctx context.Context, id uint) (*models.Student, error)
	UpdateStudent(ctx context.Context, student *models.Student) error
	DeleteStudent(ctx context.Context, id uint) error
	GetAllStudents(ctx context.Context) ([]*models.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(ctx context.Context, student *models.Student) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student.Password = string(encryptedPassword)
	if err := ValidateStudent(student); err != nil {
		return err
	}
	return s.repo.CreateStudent(ctx, student)
}
func (s *studentService) GetStudentByID(ctx context.Context, id uint) (*models.Student, error) {
	if id <= 0 {
		return nil, errors.New("ID should be Positive")
	}
	return s.repo.GetStudentByID(ctx, id)
}
func (s *studentService) UpdateStudent(ctx context.Context, student *models.Student) error {
	if student.ID <= 0 {
		return errors.New("ID should be Positive")
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student.Password = string(encryptedPassword)
	if err := ValidateStudent(student); err != nil {
		return err
	}
	return s.repo.UpdateStudent(ctx, student)
}
func (s *studentService) DeleteStudent(ctx context.Context, id uint) error {
	if id <= 0 {
		return errors.New("ID should be Positive")
	}
	return s.repo.DeleteStudent(ctx, id)
}
func (s *studentService) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	return s.repo.GetAllStudents(ctx)
}

func ValidateStudent(student *models.Student) error {
	if student.Name == "" {
		return errors.New("name is required")
	}
	if student.Age <= 0 {
		return errors.New("age must be greater than 0")
	}
	if student.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
