package repository

import (
	"context"

	"github.com/sidz111/jtw-second/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	GetStudentByID(ctx context.Context, id uint) (*models.Student, error)
	UpdateStudent(ctx context.Context, student *models.Student) error
	DeleteStudent(ctx context.Context, id uint) error
	GetAllStudents(ctx context.Context) ([]*models.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	result := r.db.WithContext(ctx).Create(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *studentRepository) GetStudentByID(ctx context.Context, id uint) (*models.Student, error) {
	var student models.Student
	result := r.db.WithContext(ctx).First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}
func (r *studentRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	result := r.db.WithContext(ctx).Save(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *studentRepository) DeleteStudent(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *studentRepository) GetAllStudents(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student
	result := r.db.WithContext(ctx).Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}
