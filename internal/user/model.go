package user

import "time"

type User struct {
	ID           int       `gorm:"column:id" json:"id" example:"1"`
	Username     string    `gorm:"column:username" json:"username" example:"johndoe"`
	PasswordHash string    `gorm:"-" json:"-"`
	Email        string    `gorm:"column:email" json:"email" example:"john.doe@example.com"`
	RoleID       int       `gorm:"column:role_id" json:"role_id" example:"1"`
	Preferences  string    `gorm:"column:preferences" json:"preferences,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
