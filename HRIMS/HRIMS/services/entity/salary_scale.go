package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type SalaryScale struct {
	ID              int32
	SalaryScaleName string
	PositionID      int32
	MinimumSalary   float64
	MaximumSalary   float64
	CurrencyType    string
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
func NewSalaryScale(name string, positionID int32, minSalary, maxSalary float64, currency string, createdBy int32) (*SalaryScale, error) {
	entity := &SalaryScale{
		SalaryScaleName: name,
		PositionID:      positionID,
		MinimumSalary:   minSalary,
		MaximumSalary:   maxSalary,
		CurrencyType:    currency,
		CreatedBy:       createdBy,
		CreatedAt:       time.Now(),
	}

	if err := entity.ValidateNewSalaryScale(); err != nil {
		log.Errorf("error validating new SalaryScale entity %v", err)
		return nil, err
	}

	return entity, nil
}
func (s *SalaryScale) ValidateNewSalaryScale() error {
	if s.SalaryScaleName == "" {
		return errors.New("error validating SalaryScale entity, salary_scale_name field required")
	}
	if s.PositionID <= 0 {
		return errors.New("error validating SalaryScale entity, position_id field required")
	}
	if s.MinimumSalary <= 0 {
		return errors.New("error validating SalaryScale entity, minimum_salary must be greater than 0")
	}
	if s.MaximumSalary <= 0 {
		return errors.New("error validating SalaryScale entity, maximum_salary must be greater than 0")
	}
	if s.MaximumSalary < s.MinimumSalary {
		return errors.New("error validating SalaryScale entity, maximum_salary must be greater than or equal to minimum_salary")
	}
	if s.CurrencyType == "" {
		return errors.New("error validating SalaryScale entity, currency_type required")
	}
	if s.CreatedBy <= 0 {
		return errors.New("error validating SalaryScale entity, createdBy field required")
	}
	return nil
}

func (s *SalaryScale) ValidateUpdateSalaryScale() error {
	if s.ID <= 0 {
		return errors.New("error validating SalaryScale entity, id field required")
	}
	if s.SalaryScaleName == "" {
		return errors.New("error validating SalaryScale entity, salary_scale_name field required")
	}
	if s.UpdatedBy <= 0 {
		return errors.New("error validating SalaryScale entity, updatedBy field required")
	}
	return nil
}
func (s *SalaryScale) MarkDeleted(userID int32) {
	s.DeletedBy = userID
	s.DeletedAt = time.Now()
}
