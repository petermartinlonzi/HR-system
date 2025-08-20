package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Department struct {
	ID           int32
	Name         string
	Description  string
	DirectorateID int32
	CreatedBy    int32
	UpdatedBy    int32
	DeletedBy    int32
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func NewDepartment(name, description string, directorateID, createdBy int32) (*Department, error) {
	entity := &Department{
		Name:          name,
		Description:   description,
		DirectorateID: directorateID,
		CreatedBy:     createdBy,
	}

	err := entity.ValidateNewDepartment()
	if err != nil {
		log.Errorf("error validating new Department entity %v", err)
		return &Department{}, err
	}

	return entity, nil
}

func (r *Department) ValidateNewDepartment() error {
	if r.Name == "" {
		return errors.New("error validating Department entity, name field required")
	}
	if r.DirectorateID <= 0 {
		return errors.New("error validating Department entity, directorateID field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Department entity, createdBy field required")
	}
	return nil
}

func (r *Department) ValidateUpdateDepartment() error {
	if r.ID <= 0 {
		return errors.New("error validating Department entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Department entity, name field required")
	}
	if r.DirectorateID <= 0 {
		return errors.New("error validating Department entity, directorateID field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Department entity, updatedBy field required")
	}
	return nil
}
