package attendance

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllLeaveRequestsByEmployeeID(ctx context.Context, employeeID int) ([]LeaveRequest, error)
	CreateLeaveRequest(ctx context.Context, leaveRequest LeaveRequest) (LeaveRequest, error)
	UpdateLeaveRequestApprovers(ctx context.Context, id int, approvers string) error
	DeleteLeaveRequest(ctx context.Context, id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllLeaveRequestsByEmployeeID(ctx context.Context, employeeID int) ([]LeaveRequest, error) {
	var leaveRequests []LeaveRequest
	if err := r.db.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&leaveRequests).Error; err != nil {
		return nil, err
	}
	return leaveRequests, nil
}

func (r *repository) CreateLeaveRequest(ctx context.Context, leaveRequest LeaveRequest) (LeaveRequest, error) {
	leaveRequest.ApplicationTime = time.Now() // 設置申請時間
	if err := r.db.WithContext(ctx).Create(&leaveRequest).Error; err != nil {
		return LeaveRequest{}, err
	}
	return leaveRequest, nil
}

func (r *repository) UpdateLeaveRequestApprovers(ctx context.Context, id int, approvers string) error {
	err := r.db.WithContext(ctx).Model(&LeaveRequest{}).
		Where("id = ?", id).
		Update("approvers", approvers).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteLeaveRequest(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&LeaveRequest{}, id).Error; err != nil {
		return err
	}
	return nil
}
