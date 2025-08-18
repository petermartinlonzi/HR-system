package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Employees struct {
	ID          int32
    FirstName   string
	LastName    string
	Email      string
	PhoneNumber int32
	Department  string
	Position    string
	HireDate    time.Time
	CreatedBy   string
	UpdatedBy   string
	DeletedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewEmployees(firstName, lastName, email string, phoneNumber int32, department, position string, hireDate time.Time, createdBy, updatedBy, deletedBy string) (*Employees, error) {
	entity := &Employees{
		FirstName:  firstName,
		LastName:   lastName,
		Email:     email,
		PhoneNumber: phoneNumber,
		Department: department,
		Position:   position,
		HireDate:   hireDate,
		CreatedBy:  createdBy,
		UpdatedBy:  updatedBy,
		DeletedBy:  deletedBy,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  time.Time{},
	}

	err := entity.ValidateNewEmployees()
	if err != nil {
		log.Errorf("error validating new Employees entity %v", err)
		return &Employees{}, err
	}

	return entity, nil
}


func (r *Employees) ValidateNewEmployees() error {
	if r.FirstName == "" {
		return errors.New("error validating Employees entity, firstName field required")
	}
	if r.CreatedBy == "" {
		return errors.New("error validating Employees entity, createdBy field required")
	}
	return nil
}

func (r *Employees) ValidateUpdateEmployees() error {
	if r.ID <= 0 {
		return errors.New("error validating Employees entity, id field required")
	}
	if r.FirstName == "" {
		return errors.New("error validating Employees entity, firstName field required")
	}
	if r.UpdatedBy == "" {
		return errors.New("error validating Employees entity, updatedBy field required")
	}
	return nil
}
