package employee

import (
	"time"
)

const (
	Active   = "active"
	Inactive = "inactive"
)

type Employee struct {
	ID           int       `json:"id" example:"1"`
	UserID       int       `json:"user_id" example:"1"`
	Name         string    `json:"name" example:"John Doe"`
	Position     string    `json:"position" example:"Software Engineer"`
	DepartmentID int       `json:"department_id" example:"1"`
	Status       string    `json:"status" example:"active"`
	ContactInfo  string    `json:"contact_info" example:"john.doe@example.com"`
	Salary       float64   `json:"salary" example:"50000"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ListEmployeesFilter struct {
	Name      string
	Position  string
	MinSalary float64
	MaxSalary float64
}
