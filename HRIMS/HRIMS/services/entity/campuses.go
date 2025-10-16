package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Campuses struct {
	ID          int32
	Name        string
	Description string
	Resvd1      string
	Resvd2      string
	Resvd3      string
	Resvd4      string
	Resvd5      string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

// Constructor
func NewCampuses(name, description string, createdBy int32) (*Campuses, error) {
	entity := &Campuses{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewCampuses()
	if err != nil {
		log.Errorf("error validating new Campuses entity %v", err)
		return &Campuses{}, err
	}

	return entity, nil
}

// Validation wakati wa ku-create campuses
func (r *Campuses) ValidateNewCampuses() error {
	if r.Name == "" {
		return errors.New("error validating Campuses entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Campuses entity, createdBy field required")
	}
	return nil
}

// Validation wakati wa ku-update campuses
func (r *Campuses) ValidateUpdateCampuses() error {
	if r.ID <= 0 {
		return errors.New("error validating Campuses entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Campuses entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Campuses entity, updatedBy field required")
	}
	return nil
}
