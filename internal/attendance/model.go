package attendance

import (
	"time"
)

type LeaveRequest struct {
	ID              int       `gorm:"column:id;primaryKey" json:"id" description:"單號"`
	EmployeeID      int       `gorm:"column:employee_id" json:"employee_id" description:"員工編號"`
	LeaveType       string    `gorm:"column:leave_type" json:"leave_type" description:"假別"`
	StartTime       time.Time `gorm:"column:start_time" json:"start_time" description:"開始時間"`
	EndTime         time.Time `gorm:"column:end_time" json:"end_time" description:"結束時間"`
	TotalHours      float64   `gorm:"column:total_hours" json:"total_hours" description:"總時數"`
	SubstituteID    int       `gorm:"column:substitute_id" json:"substitute_id" description:"職務代理人"`
	Description     string    `gorm:"column:description" json:"description" description:"申請說明"`
	Approvers       string    `gorm:"column:approvers" json:"approvers" description:"該單審核鏈有哪些人/角色"`
	ApplicationTime time.Time `gorm:"column:application_time" json:"application_time" description:"申請時間"`
}

type Attendance struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	CheckIn    time.Time
	CheckOut   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
