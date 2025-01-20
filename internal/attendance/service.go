package attendance

import (
	"context"
	"fmt"

	"hr-system/internal/user"
)

type Service interface {
	GetLeaveRequests(ctx context.Context, employeeID int) ([]LeaveRequest, error)
	SubmitLeaveRequest(ctx context.Context, leaveRequest LeaveRequest) (LeaveRequest, error)
	ApproveOrRejectLeaveRequest(ctx context.Context, id int, status string) error
	DeleteLeaveRequest(ctx context.Context, id int) error
}

type service struct {
	repo     Repository
	userRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) Service {
	return &service{repo: repo, userRepo: userRepo}
}

func (s *service) GetLeaveRequests(ctx context.Context, employeeID int) ([]LeaveRequest, error) {
	return s.repo.GetAllLeaveRequestsByEmployeeID(ctx, employeeID)
}

func (s *service) SubmitLeaveRequest(ctx context.Context, leaveRequest LeaveRequest) (LeaveRequest, error) {

	// TODO: generate approvers
	approvers := fmt.Sprintf("1:1:1737395670,2:0:0,3:0:0")
	leaveRequest.Approvers = approvers

	return s.repo.CreateLeaveRequest(ctx, leaveRequest)
}

func (s *service) ApproveOrRejectLeaveRequest(ctx context.Context, id int, status string) error {

	// TODO: combine approvers
	approvers := fmt.Sprintf("1:1:1737395670,2:1:1737395670,3:0:0")

	return s.repo.UpdateLeaveRequestApprovers(ctx, id, approvers)
}

func (s *service) DeleteLeaveRequest(ctx context.Context, id int) error {
	return s.repo.DeleteLeaveRequest(ctx, id)
}
