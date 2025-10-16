package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type LeaveBalance struct {
	ID             int32
	UserID         int32
	LeaveType      string
	Year           int32
	TotalEntitled  int32
	UsedDays       int32
	RemainingDays  int32
	CreatedBy      int32
	UpdatedBy      int32
	DeletedBy      int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func NewLeaveBalance(userID int32, leaveType string, year, totalEntitled, createdBy int32) (*LeaveBalance, error) {
	entity := &LeaveBalance{
		UserID:        userID,
		LeaveType:     leaveType,
		Year:          year,
		TotalEntitled: totalEntitled,
		UsedDays:      0,
		RemainingDays: totalEntitled,
		CreatedBy:     createdBy,
	}

	err := entity.ValidateNewLeaveBalance()
	if err != nil {
		log.Errorf("error validating new LeaveBalance entity %v", err)
		return &LeaveBalance{}, err
	}

	return entity, nil
}

func (r *LeaveBalance) ValidateNewLeaveBalance() error {
	if r.UserID <= 0 {
		return errors.New("error validating LeaveBalance entity, userID required")
	}
	if r.LeaveType == "" {
		return errors.New("error validating LeaveBalance entity, leaveType required")
	}
	if r.Year <= 0 {
		return errors.New("error validating LeaveBalance entity, year required")
	}
	if r.TotalEntitled <= 0 {
		return errors.New("error validating LeaveBalance entity, totalEntitled must be greater than 0")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating LeaveBalance entity, createdBy required")
	}
	return nil
}

func (r *LeaveBalance) ValidateUpdateLeaveBalance() error {
	if r.ID <= 0 {
		return errors.New("error validating LeaveBalance entity, id required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating LeaveBalance entity, updatedBy required")
	}
	return nil
}
