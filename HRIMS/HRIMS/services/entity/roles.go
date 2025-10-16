package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Role struct {
	ID          int32
	Name        string
	Description string
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewRole(name, description string, createdBy int32) (*Role, error) {
	entity := &Role{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
	}

	if err := entity.ValidateNewRole(); err != nil {
		log.Errorf("error validating new Role entity %v", err)
		return nil, err
	}

	return entity, nil
}
func (r *Role) ValidateNewRole() error {
	if r.Name == "" {
		return errors.New("error validating Role entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Role entity, createdBy field required")
	}
	return nil
}
func (r *Role) ValidateUpdateRole() error {
	if r.ID <= 0 {
		return errors.New("error validating Role entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Role entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Role entity, updatedBy field required")
	}
	return nil
}

// Soft delete helper
func (r *Role) MarkDeleted(userID int32) {
	r.DeletedBy = userID
	r.DeletedAt = time.Now()
}
