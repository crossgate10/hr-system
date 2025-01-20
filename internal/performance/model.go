package performance

import "time"

type Performance struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Rating     int // 1 to 5 or any rating scale
	Feedback   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Repository interface {
	CreatePerformance(performance *Performance) error
	GetPerformanceByEmployeeID(employeeID uint) (*Performance, error)
	GetAllPerformances() ([]Performance, error)
	UpdatePerformance(performance *Performance) error
	DeletePerformance(id uint) error
}

type Service interface {
	EvaluatePerformance(employeeID uint, rating int, feedback string) (*Performance, error)
	GetPerformanceDetails(employeeID uint) (*Performance, error)
	UpdatePerformanceDetails(performance *Performance) error
	RemovePerformance(id uint) error
}
