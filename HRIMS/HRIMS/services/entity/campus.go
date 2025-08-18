package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Campus struct {
	CampusID     int32
	CampusName   string
	Description  string
	DepartmentID int32
	CreatedBy    int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

// Constructor
func NewCampus(campusName string, description string, departmentID int32, createdBy int32) (*Campus, error) {
	entity := &Campus{
		CampusName:   campusName,
		Description:  description,
		DepartmentID: departmentID,
		CreatedBy:    createdBy,
	}

	err := entity.ValidateNewCampus()
	if err != nil {
		log.Errorf("error validating new Campus entity %v", err)
		return &Campus{}, err
	}

	return entity, nil
}

// Validation wakati wa ku-create campus
func (r *Campus) ValidateNewCampus() error {
	if r.CampusName == "" {
		return errors.New("error validating Campus entity, campusName field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Campus entity, createdBy field required")
	}
	return nil
}

// Validation wakati wa ku-update campus
func (r *Campus) ValidateUpdateCampus() error {
	if r.CampusID <= 0 {
		return errors.New("error validating Campus entity, campusID field required")
	}
	if r.CampusName == "" {
		return errors.New("error validating Campus entity, campusName field required")
	}
	if r.DepartmentID <= 0 {
		return errors.New("error validating Campus entity, departmentID field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Campus entity, createdBy field required")
	}
	return nil
}
