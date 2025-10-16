package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Approval struct {
	ID         int32
	RequestID  int32
	ApprovedBy int32
	Status     string
	Comment    string
	ApprovedAt time.Time
	CreatedBy  int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func NewApproval(requestID, approvedBy, createdBy int32, status, comment string) (*Approval, error) {
	entity := &Approval{
		RequestID:  requestID,
		ApprovedBy: approvedBy,
		Status:     status,
		Comment:    comment,
		CreatedBy:  createdBy,
	}

	err := entity.ValidateNewApproval()
	if err != nil {
		log.Errorf("error validating new Approval entity %v", err)
		return &Approval{}, err
	}

	return entity, nil
}

func (r *Approval) ValidateNewApproval() error {
	if r.RequestID <= 0 {
		return errors.New("error validating Approval entity, requestID field required")
	}
	if r.ApprovedBy <= 0 {
		return errors.New("error validating Approval entity, approvedBy field required")
	}
	if r.Status == "" {
		return errors.New("error validating Approval entity, status field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Approval entity, createdBy field required")
	}
	return nil
}

func (r *Approval) ValidateUpdateApproval() error {
	if r.ID <= 0 {
		return errors.New("error validating Approval entity, id field required")
	}
	if r.Status == "" {
		return errors.New("error validating Approval entity, status field required")
	}
	if r.ApprovedBy <= 0 {
		return errors.New("error validating Approval entity, approvedBy field required")
	}
	return nil
}
