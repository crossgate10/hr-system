package training

import "time"

type Training struct {
	ID               uint `gorm:"primaryKey"`
	EmployeeID       uint
	CourseName       string
	CompletionStatus string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Repository interface {
	CreateTraining(training *Training) error
	GetTrainingByEmployeeID(employeeID uint) (*Training, error)
	GetAllTrainings() ([]Training, error)
	UpdateTraining(training *Training) error
	DeleteTraining(id uint) error
}

type Service interface {
	EnrollTraining(employeeID uint, courseName, completionStatus string) (*Training, error)
	GetTrainingDetails(employeeID uint) (*Training, error)
	UpdateTrainingDetails(training *Training) error
	RemoveTraining(id uint) error
}
