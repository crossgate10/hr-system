package user

import "time"

// TODO: use RBAC

type Approver struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Seq  int    `gorm:"column:seq" json:"seq"`
}

type Role struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type Department struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type RoleDepartment struct {
	RoleID       int `gorm:"column:role_id" json:"role_id"`
	DepartmentID int `gorm:"column:department_id" json:"department_id"`
	ApproverID   int `gorm:"column:approver_id" json:"approver_id"`
}

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
