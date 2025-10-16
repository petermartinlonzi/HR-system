package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type ContractVersion struct {
	ID             int32
	ContractID     int32
	VersionNumber  int32
	Salary         float64
	Benefits       string
	WorkingHours   string
	ProbationPeriod string
	SignedBy       int32
	SignedAt       time.Time
	CreatedBy      int32
	UpdatedBy      int32
	DeletedBy      int32
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
}

// NewContractVersion creates a new ContractVersion entity
func NewContractVersion(contractID int32, versionNumber int32, salary float64, benefits, workingHours, probationPeriod string, signedBy int32, createdBy int32) (*ContractVersion, error) {
	entity := &ContractVersion{
		ContractID:     contractID,
		VersionNumber:  versionNumber,
		Salary:         salary,
		Benefits:       benefits,
		WorkingHours:   workingHours,
		ProbationPeriod: probationPeriod,
		SignedBy:       signedBy,
		CreatedBy:      createdBy,
		SignedAt:       time.Now(),
	}

	err := entity.ValidateNewContractVersion()
	if err != nil {
		log.Errorf("error validating new ContractVersion entity %v", err)
		return &ContractVersion{}, err
	}

	return entity, nil
}

// ValidateNewContractVersion validates required fields for creation
func (r *ContractVersion) ValidateNewContractVersion() error {
	if r.ContractID <= 0 {
		return errors.New("error validating ContractVersion entity, contractID field required")
	}
	if r.VersionNumber <= 0 {
		return errors.New("error validating ContractVersion entity, versionNumber field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating ContractVersion entity, createdBy field required")
	}
	return nil
}

// ValidateUpdateContractVersion validates required fields for update
func (r *ContractVersion) ValidateUpdateContractVersion() error {
	if r.ID <= 0 {
		return errors.New("error validating ContractVersion entity, id field required")
	}
	if r.ContractID <= 0 {
		return errors.New("error validating ContractVersion entity, contractID field required")
	}
	if r.VersionNumber <= 0 {
		return errors.New("error validating ContractVersion entity, versionNumber field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating ContractVersion entity, updatedBy field required")
	}
	return nil
}
