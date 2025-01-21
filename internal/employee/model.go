package employee

import (
	"time"
)

const (
	Active   = "active"
	Inactive = "inactive"
)

type Employee struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"John Doe"`
	Title       string    `json:"title" example:"Software Engineer"`
	Role        string    `json:"role" example:"backend_member"`
	ContactInfo string    `json:"contact_info" example:"john.doe@example.com"`
	Salary      float64   `json:"salary" example:"50000"`
	Status      string    `json:"status" example:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListEmployeesFilter struct {
	Name      string
	Title     string
	MinSalary float64
	MaxSalary float64
}
