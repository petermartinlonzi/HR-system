package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Officer struct {
	ID           int32
	UserID       int32
	DepartmentID int32
	Position     string
	Phone        string
	Designation  string
	CreatedBy    int32
	UpdatedBy    int32
	DeletedBy    int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

// Constructor for creating a new Officer entity
func NewOfficer(userID, departmentID int32, position, phone, designation string, createdBy int32) (*Officer, error) {
	entity := &Officer{
		UserID:       userID,
		DepartmentID: departmentID,
		Position:     position,
		Phone:        phone,
		Designation:  designation,
		CreatedBy:    createdBy,
	}

	err := entity.ValidateNewOfficer()
	if err != nil {
		log.Errorf("error validating new Officer entity %v", err)
		return &Officer{}, err
	}

	return entity, nil
}

// Validation for creating a new Officer
func (r *Officer) ValidateNewOfficer() error {
	if r.UserID <= 0 {
		return errors.New("error validating Officer entity, userID field required")
	}
	if r.DepartmentID <= 0 {
		return errors.New("error validating Officer entity, departmentID field required")
	}
	if r.Position == "" {
		return errors.New("error validating Officer entity, position field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Officer entity, createdBy field required")
	}
	return nil
}

// Validation for updating an existing Officer
func (r *Officer) ValidateUpdateOfficer() error {
	if r.ID <= 0 {
		return errors.New("error validating Officer entity, id field required")
	}
	if r.Position == "" {
		return errors.New("error validating Officer entity, position field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Officer entity, updatedBy field required")
	}
	return nil
}
