package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type OfficerRole struct {
	ID          int32
	RoleName    string
	Description string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// Constructor for creating a new OfficerRole entity
func NewOfficerRole(roleName, description string, createdBy int32) (*OfficerRole, error) {
	entity := &OfficerRole{
		RoleName:    roleName,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewOfficerRole()
	if err != nil {
		log.Errorf("error validating new OfficerRole entity %v", err)
		return &OfficerRole{}, err
	}

	return entity, nil
}

// Validation for creating a new OfficerRole
func (r *OfficerRole) ValidateNewOfficerRole() error {
	if r.RoleName == "" {
		return errors.New("error validating OfficerRole entity, roleName field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating OfficerRole entity, createdBy field required")
	}
	return nil
}

// Validation for updating an existing OfficerRole
func (r *OfficerRole) ValidateUpdateOfficerRole() error {
	if r.ID <= 0 {
		return errors.New("error validating OfficerRole entity, id field required")
	}
	if r.RoleName == "" {
		return errors.New("error validating OfficerRole entity, roleName field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating OfficerRole entity, updatedBy field required")
	}
	return nil
}
