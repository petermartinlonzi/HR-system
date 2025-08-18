package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Departments struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewDepartments(name, description string, createdBy int32) (*Departments, error) {
	entity := &Departments{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewDepartments()
	if err != nil {
		log.Errorf("error validating new Departments entity %v", err)
		return &Departments{}, err
	}

	return entity, nil
}

func (r *Departments) ValidateNewDepartments() error {
	if r.Name == "" {
		return errors.New("error validating Departments entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Departments entity, createdBy field required")
	}
	return nil
}

func (r *Departments) ValidateUpdateDepartments() error {
	if r.ID <= 0 {
		return errors.New("error validating Departments entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Departments entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Departments entity, updatedBy field required")
	}
	return nil
}
