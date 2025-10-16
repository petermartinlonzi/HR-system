package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Contract struct {
	ID               int32
	ApplicationID    int32
	EmployeeID       int32
	JobID            int32
	ContractType     string
	StartDate        time.Time
	EndDate          *time.Time
	SignedAt         time.Time
	ContractStatusID int32
	Resvd1    string
	Resvd2    string
	Resvd3    string
	Resvd4    string
	Resvd5    string
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Constructor
func NewContract(applicationID, employeeID, jobID int32, contractType string, startDate time.Time, endDate *time.Time, createdBy int32) (*Contract, error) {
	entity := &Contract{
		ApplicationID:    applicationID,
		EmployeeID:       employeeID,
		JobID:            jobID,
		ContractType:     contractType,
		StartDate:        startDate,
		EndDate:          endDate,
		ContractStatusID: 1, // default value
		CreatedBy:        createdBy,
		SignedAt:         time.Now(),
	}

	err := entity.ValidateNewContract()
	if err != nil {
		log.Errorf("error validating new Contract entity %v", err)
		return &Contract{}, err
	}

	return entity, nil
}

// Validation wakati wa ku-create contract
func (r *Contract) ValidateNewContract() error {
	if r.ApplicationID <= 0 {
		return errors.New("error validating Contract entity, applicationID field required")
	}
	if r.EmployeeID <= 0 {
		return errors.New("error validating Contract entity, employeeID field required")
	}
	if r.JobID <= 0 {
		return errors.New("error validating Contract entity, jobID field required")
	}
	if r.StartDate.IsZero() {
		return errors.New("error validating Contract entity, startDate field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Contract entity, createdBy field required")
	}
	return nil
}

// Validation wakati wa ku-update contract
func (r *Contract) ValidateUpdateContract() error {
	if r.ID <= 0 {
		return errors.New("error validating Contract entity, id field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Contract entity, updatedBy field required")
	}
	if r.ContractStatusID <= 0 {
		return errors.New("error validating Contract entity, contractStatusID field required")
	}
	return nil
}
