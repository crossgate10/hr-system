package notification

import "time"

type Notification struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	Message    string
	ReadStatus bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Repository interface {
	CreateNotification(notification *Notification) error
	GetNotificationByEmployeeID(employeeID uint) ([]Notification, error)
	GetAllNotifications() ([]Notification, error)
	UpdateNotification(notification *Notification) error
	DeleteNotification(id uint) error
}

type Service interface {
	SendNotification(employeeID uint, message string) (*Notification, error)
	GetNotificationsForEmployee(employeeID uint) ([]Notification, error)
	MarkNotificationAsRead(id uint) error
	RemoveNotification(id uint) error
}
