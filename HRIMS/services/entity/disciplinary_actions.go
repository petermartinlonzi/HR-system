package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type DisciplinaryAction struct {
	ID          int32
	EmployeeID  int32
	ReportedBy  int32
	IncidentDate time.Time
	Description string
	ActionTaken string
	ActionDate  *time.Time
	Duration    *int32
	Status      string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

// NewDisciplinaryAction creates a new DisciplinaryAction entity
func NewDisciplinaryAction(employeeID, reportedBy int32, incidentDate time.Time, description string, createdBy int32) (*DisciplinaryAction, error) {
	entity := &DisciplinaryAction{
		EmployeeID:  employeeID,
		ReportedBy:  reportedBy,
		IncidentDate: incidentDate,
		Description: description,
		Status:      "Resolved",
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewDisciplinaryAction()
	if err != nil {
		log.Errorf("error validating new DisciplinaryAction entity %v", err)
		return &DisciplinaryAction{}, err
	}

	return entity, nil
}

// ValidateNewDisciplinaryAction validates required fields when creating a new record
func (r *DisciplinaryAction) ValidateNewDisciplinaryAction() error {
	if r.EmployeeID <= 0 {
		return errors.New("error validating DisciplinaryAction, employeeID required")
	}
	if r.ReportedBy <= 0 {
		return errors.New("error validating DisciplinaryAction, reportedBy required")
	}
	if r.IncidentDate.IsZero() {
		return errors.New("error validating DisciplinaryAction, incidentDate required")
	}
	if r.Description == "" {
		return errors.New("error validating DisciplinaryAction, description required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating DisciplinaryAction, createdBy required")
	}
	return nil
}

// ValidateUpdateDisciplinaryAction validates required fields when updating a record
func (r *DisciplinaryAction) ValidateUpdateDisciplinaryAction() error {
	if r.ID <= 0 {
		return errors.New("error validating DisciplinaryAction, id required")
	}
	if r.EmployeeID <= 0 {
		return errors.New("error validating DisciplinaryAction, employeeID required")
	}
	if r.Description == "" {
		return errors.New("error validating DisciplinaryAction, description required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating DisciplinaryAction, updatedBy required")
	}
	return nil
}
