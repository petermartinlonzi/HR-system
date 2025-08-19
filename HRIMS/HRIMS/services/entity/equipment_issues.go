package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type EquipmentIssue struct {
	ID                int32
	EquipmentID       int32
	IssuedTo          int32
	IssueDate         time.Time
	ReturnDate        *time.Time
	ReturnedCondition string
	CreatedBy         int32
	UpdatedBy         int32
	DeletedBy         int32
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

func NewEquipmentIssue(equipmentID, issuedTo, createdBy int32, issueDate time.Time) (*EquipmentIssue, error) {
	entity := &EquipmentIssue{
		EquipmentID: equipmentID,
		IssuedTo:    issuedTo,
		IssueDate:   issueDate,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewEquipmentIssue()
	if err != nil {
		log.Errorf("error validating new EquipmentIssue entity %v", err)
		return &EquipmentIssue{}, err
	}

	return entity, nil
}

func (r *EquipmentIssue) ValidateNewEquipmentIssue() error {
	if r.EquipmentID <= 0 {
		return errors.New("error validating EquipmentIssue entity, equipmentID required")
	}
	if r.IssuedTo <= 0 {
		return errors.New("error validating EquipmentIssue entity, issuedTo required")
	}
	if r.IssueDate.IsZero() {
		return errors.New("error validating EquipmentIssue entity, issueDate required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating EquipmentIssue entity, createdBy required")
	}
	return nil
}

func (r *EquipmentIssue) ValidateUpdateEquipmentIssue() error {
	if r.ID <= 0 {
		return errors.New("error validating EquipmentIssue entity, id required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating EquipmentIssue entity, updatedBy required")
	}
	return nil
}
