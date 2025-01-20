package payroll

import "time"

type Payroll struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Salary     float64
	Bonus      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Repository interface {
	CreatePayroll(payroll *Payroll) error
	GetPayrollByEmployeeID(employeeID uint) (*Payroll, error)
	GetAllPayrolls() ([]Payroll, error)
	UpdatePayroll(payroll *Payroll) error
	DeletePayroll(id uint) error
}

type Service interface {
	ProcessPayroll(employeeID uint, salary, bonus float64) (*Payroll, error)
	GetPayrollDetails(employeeID uint) (*Payroll, error)
	UpdatePayrollDetails(payroll *Payroll) error
	RemovePayroll(id uint) error
}
