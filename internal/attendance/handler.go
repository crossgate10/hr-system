package attendance

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, service Service) {
	h := &Handler{service: service}
	router.GET("/leave-requests", h.ListLeaveRequests)
	router.POST("/leave-requests", h.SubmitLeaveRequest)
	router.PUT("/leave-requests/:id", h.UpdateLeaveRequest)
	router.DELETE("/leave-requests/:id", h.DeleteLeaveRequest)
}

type Handler struct {
	service Service
}

func (h *Handler) ListLeaveRequests(c *gin.Context) {
	// TODO: get user id from token
	userID := 1

	leaveRequests, err := h.service.GetLeaveRequests(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leaveRequests)
}

type SubmitLeaveRequest struct {
	EmployeeID   int    `json:"employee_id"`
	StartDate    int64  `json:"start_date"`
	EndDate      int64  `json:"end_date"`
	LeaveType    string `json:"leave_type"`
	SubstituteID int    `json:"substitute_id"`
	Reason       string `json:"reason"`
}

type LeaveRequestResponse struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	StartDate  int64  `json:"start_date"`
	EndDate    int64  `json:"end_date"`
	Reason     string `json:"reason"`
	Approvers  string `json:"approvers" description:"id1:status:unix1,id2:status:unix2" example:"1:1:1737393916,2:0:0"`
}

func (h *Handler) SubmitLeaveRequest(c *gin.Context) {

	// TODO: check role

	var req SubmitLeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leaveRequest := LeaveRequest{
		EmployeeID:   req.EmployeeID,
		LeaveType:    req.LeaveType,
		StartTime:    time.Unix(req.StartDate, 0),
		EndTime:      time.Unix(req.EndDate, 0),
		SubstituteID: req.SubstituteID,
		Description:  req.Reason,
	}
	createdLeaveRequest, err := h.service.SubmitLeaveRequest(c.Request.Context(), leaveRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := LeaveRequestResponse{
		ID:         createdLeaveRequest.ID,
		EmployeeID: createdLeaveRequest.EmployeeID,
		StartDate:  createdLeaveRequest.StartTime.Unix(),
		EndDate:    createdLeaveRequest.EndTime.Unix(),
		Reason:     createdLeaveRequest.Description,
		Approvers:  createdLeaveRequest.Approvers,
	}

	c.JSON(http.StatusCreated, resp)
}

type UpdateLeaveRequest struct {
	Status string `json:"status"`
}

func (h *Handler) UpdateLeaveRequest(c *gin.Context) {

	// TODO: check role
	// TODO: get approver id from token

	leaveRequestID, _ := strconv.Atoi(c.Param("id"))
	if leaveRequestID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req UpdateLeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ApproveOrRejectLeaveRequest(c.Request.Context(), leaveRequestID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave request updated successfully"})
}

func (h *Handler) DeleteLeaveRequest(c *gin.Context) {

	// TODO: check role
	// TODO: consider soft delete

	leaveRequestID, _ := strconv.Atoi(c.Param("id"))
	if leaveRequestID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	err := h.service.DeleteLeaveRequest(c.Request.Context(), leaveRequestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave request deleted successfully"})
}
