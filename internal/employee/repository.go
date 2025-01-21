package employee

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	CreateEmployee(ctx context.Context, employee *Employee) error
	GetEmployeeByID(ctx context.Context, id int) (*Employee, error)
	ListEmployees(ctx context.Context, filter *ListEmployeesFilter) ([]Employee, error)
	UpdateEmployee(ctx context.Context, employee *Employee) error
	DeleteEmployee(ctx context.Context, id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateEmployee(ctx context.Context, employee *Employee) error {
	return r.db.WithContext(ctx).Create(employee).Error
}

func (r *repository) GetEmployeeByID(ctx context.Context, id int) (*Employee, error) {
	var employee Employee
	if err := r.db.WithContext(ctx).First(&employee, id).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *repository) ListEmployees(ctx context.Context, filter *ListEmployeesFilter) ([]Employee, error) {
	var employees []Employee
	query := r.db.WithContext(ctx)

	if filter != nil {
		if filter.Name != "" {
			query = query.Where("name LIKE ?", "%"+filter.Name+"%")
		}
		if filter.Title != "" {
			query = query.Where("title LIKE ?", "%"+filter.Title+"%")
		}
		if filter.MinSalary > 0 {
			query = query.Where("salary >= ?", filter.MinSalary)
		}
		if filter.MaxSalary > 0 {
			query = query.Where("salary <= ?", filter.MaxSalary)
		}
	}

	if err := query.Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *repository) UpdateEmployee(ctx context.Context, employee *Employee) error {
	return r.db.WithContext(ctx).Save(employee).Error
}

func (r *repository) DeleteEmployee(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&Employee{}, id).Error
}
