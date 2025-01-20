package compliance

import "time"

type Compliance struct {
	ID            uint `gorm:"primaryKey"`
	PolicyName    string
	Description   string
	EffectiveDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Repository interface {
	CreateCompliance(compliance *Compliance) error
	GetComplianceByID(id uint) (*Compliance, error)
	GetAllCompliances() ([]Compliance, error)
	UpdateCompliance(compliance *Compliance) error
	DeleteCompliance(id uint) error
}

type Service interface {
	AddCompliance(policyName, description string, effectiveDate time.Time) (*Compliance, error)
	GetComplianceDetails(id uint) (*Compliance, error)
	UpdateComplianceDetails(compliance *Compliance) error
	RemoveCompliance(id uint) error
}
