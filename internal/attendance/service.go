package attendance

import (
	"context"
	"fmt"
)

const (
	backendMember   = "backend_member"
	frontendMember  = "frontend_member"
	pmMember        = "pm_member"
	qaMember        = "qa_member"
	backendManager  = "backend_manager"
	frontendManager = "frontend_manager"
	pmManager       = "pm_manager"
	qaManager       = "qa_manager"
	hr              = "hr"
	hrManager       = "hr_manager"
	boss            = "boss"
)

var (
	approvalRules = map[string][]string{
		backendMember:   {hr, backendMember, backendManager},
		frontendMember:  {hr, frontendMember, frontendManager},
		pmMember:        {hr, pmMember, pmManager},
		qaMember:        {hr, qaMember, qaManager},
		backendManager:  {hr, backendMember},
		frontendManager: {hr, frontendMember},
		pmManager:       {hr, pmMember},
		qaManager:       {hr, qaMember},
		hr:              {hrManager},
		hrManager:       {hr},
		boss:            {},
	}
)

type Service interface {
	GetLeaveRequests(ctx context.Context, employeeID int) ([]LeaveRequest, error)
	SubmitLeaveRequest(ctx context.Context, leaveRequest LeaveRequest) (LeaveRequest, error)
	ApproveOrRejectLeaveRequest(ctx context.Context, id int, status string) error
	DeleteLeaveRequest(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
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
