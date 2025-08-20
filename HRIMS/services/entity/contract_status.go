package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type ContractStatus struct {
	ID         int32
	StatusName string
	Description string
	CreatedBy  int32
	UpdatedBy  int32
	DeletedBy  int32
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

func NewContractStatus(statusName, description string, createdBy int32) (*ContractStatus, error) {
	entity := &ContractStatus{
		StatusName:  statusName,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewContractStatus()
	if err != nil {
		log.Errorf("error validating new ContractStatus entity %v", err)
		return &ContractStatus{}, err
	}

	return entity, nil
}

func (r *ContractStatus) ValidateNewContractStatus() error {
	if r.StatusName == "" {
		return errors.New("error validating ContractStatus entity, statusName field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating ContractStatus entity, createdBy field required")
	}
	return nil
}

func (r *ContractStatus) ValidateUpdateContractStatus() error {
	if r.ID <= 0 {
		return errors.New("error validating ContractStatus entity, id field required")
	}
	if r.StatusName == "" {
		return errors.New("error validating ContractStatus entity, statusName field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating ContractStatus entity, updatedBy field required")
	}
	return nil
}
