package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type LeaveRequest struct {
	ID         int32
	UserID     int32
	LeaveType  string
	Reason     string
	StartDate  time.Time
	EndDate    time.Time
	TotalDays  int32
	Status     string
	AppliedAt  time.Time
	ReviewedBy int32
	ReviewedAt time.Time
	Comment    string
	CreatedBy  int32
	UpdatedBy  int32
	DeletedBy  int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

// Constructor for creating a new LeaveRequest entity
func NewLeaveRequest(userID int32, leaveType, reason string, startDate, endDate time.Time, reviewedBy, createdBy int32) (*LeaveRequest, error) {
	if endDate.Before(startDate) {
		return nil, errors.New("endDate cannot be before startDate")
	}

	totalDays := int32(endDate.Sub(startDate).Hours()/24) + 1

	entity := &LeaveRequest{
		UserID:     userID,
		LeaveType:  leaveType,
		Reason:     reason,
		StartDate:  startDate,
		EndDate:    endDate,
		TotalDays:  totalDays,
		Status:     "Pending",
		AppliedAt:  time.Now(),
		ReviewedBy: reviewedBy,
		CreatedBy:  createdBy,
	}

	err := entity.ValidateNewLeaveRequest()
	if err != nil {
		log.Errorf("error validating new LeaveRequest entity %v", err)
		return &LeaveRequest{}, err
	}

	return entity, nil
}

// Validation for creating a new LeaveRequest
func (r *LeaveRequest) ValidateNewLeaveRequest() error {
	if r.UserID <= 0 {
		return errors.New("error validating LeaveRequest entity, userID field required")
	}
	if r.LeaveType == "" {
		return errors.New("error validating LeaveRequest entity, leaveType field required")
	}
	if r.StartDate.IsZero() || r.EndDate.IsZero() {
		return errors.New("error validating LeaveRequest entity, startDate and endDate fields required")
	}
	if r.ReviewedBy <= 0 {
		return errors.New("error validating LeaveRequest entity, reviewedBy field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating LeaveRequest entity, createdBy field required")
	}
	return nil
}

// Validation for updating an existing LeaveRequest
func (r *LeaveRequest) ValidateUpdateLeaveRequest() error {
	if r.ID <= 0 {
		return errors.New("error validating LeaveRequest entity, id field required")
	}
	if r.LeaveType == "" {
		return errors.New("error validating LeaveRequest entity, leaveType field required")
	}
	if r.StartDate.IsZero() || r.EndDate.IsZero() {
		return errors.New("error validating LeaveRequest entity, startDate and endDate fields required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating LeaveRequest entity, updatedBy field required")
	}
	return nil
}
