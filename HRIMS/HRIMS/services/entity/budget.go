package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Budget struct {
	BudgetID        int32
	RequestID       int32
	SalaryScaleID   int32
	NumberOfOfficers int32
	TotalBudget     float64
	Description     string
	CreatedBy       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

// Constructor
func NewBudget(requestID, salaryScaleID, numberOfOfficers int32, totalBudget float64, createdBy string, description string) (*Budget, error) {
	entity := &Budget{
		RequestID:        requestID,
		SalaryScaleID:    salaryScaleID,
		NumberOfOfficers: numberOfOfficers,
		TotalBudget:      totalBudget,
		Description:      description,
		CreatedBy:        createdBy,
	}

	err := entity.ValidateNewBudget()
	if err != nil {
		log.Errorf("error validating new Budget entity %v", err)
		return &Budget{}, err
	}

	return entity, nil
}

// Validation wakati wa ku-create budget
func (r *Budget) ValidateNewBudget() error {
	if r.RequestID <= 0 {
		return errors.New("error validating Budget entity, requestID field required")
	}
	if r.SalaryScaleID <= 0 {
		return errors.New("error validating Budget entity, salaryScaleID field required")
	}
	if r.NumberOfOfficers <= 0 {
		return errors.New("error validating Budget entity, numberOfOfficers must be greater than zero")
	}
	if r.TotalBudget <= 0 {
		return errors.New("error validating Budget entity, totalBudget must be greater than zero")
	}
	if r.CreatedBy == "" {
		return errors.New("error validating Budget entity, createdBy field required")
	}
	return nil
}

// Validation wakati wa ku-update budget
func (r *Budget) ValidateUpdateBudget() error {
	if r.BudgetID <= 0 {
		return errors.New("error validating Budget entity, budgetID field required")
	}
	if r.TotalBudget <= 0 {
		return errors.New("error validating Budget entity, totalBudget must be greater than zero")
	}
	return nil
}
