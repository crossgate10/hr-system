package recruitment

import "time"

type Candidate struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	PositionApplied string
	ContactInfo     string
	Status          string // applied, interviewed, hired, etc.
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Repository interface {
	CreateCandidate(candidate *Candidate) error
	GetCandidateByID(id uint) (*Candidate, error)
	GetAllCandidates() ([]Candidate, error)
	UpdateCandidate(candidate *Candidate) error
	DeleteCandidate(id uint) error
}

type Service interface {
	AddCandidate(name, positionApplied, contactInfo string) (*Candidate, error)
	GetCandidateDetails(id uint) (*Candidate, error)
	UpdateCandidateDetails(candidate *Candidate) error
	RemoveCandidate(id uint) error
}
