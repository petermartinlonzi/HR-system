package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Driver struct {
	DriverID          int32
	FirstName         string
	LastName          string
	LicenseNumber     string
	LicenseExpiry     *time.Time
	PhoneNumber       string
	EmploymentStatus  string
	AssignedVehicleID int32
	CreatedBy         int32
	UpdatedBy         int32
	DeletedBy         int32
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
}

func NewDriver(firstName, lastName, licenseNumber string, createdBy int32) (*Driver, error) {
	entity := &Driver{
		FirstName:        firstName,
		LastName:         lastName,
		LicenseNumber:    licenseNumber,
		EmploymentStatus: "Active",
		CreatedBy:        createdBy,
	}

	err := entity.ValidateNewDriver()
	if err != nil {
		log.Errorf("error validating new Driver entity %v", err)
		return &Driver{}, err
	}

	return entity, nil
}

func (r *Driver) ValidateNewDriver() error {
	if r.FirstName == "" {
		return errors.New("error validating Driver entity, firstName field required")
	}
	if r.LastName == "" {
		return errors.New("error validating Driver entity, lastName field required")
	}
	if r.LicenseNumber == "" {
		return errors.New("error validating Driver entity, licenseNumber field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Driver entity, createdBy field required")
	}
	return nil
}

func (r *Driver) ValidateUpdateDriver() error {
	if r.DriverID <= 0 {
		return errors.New("error validating Driver entity, driverID field required")
	}
	if r.FirstName == "" {
		return errors.New("error validating Driver entity, firstName field required")
	}
	if r.LastName == "" {
		return errors.New("error validating Driver entity, lastName field required")
	}
	if r.LicenseNumber == "" {
		return errors.New("error validating Driver entity, licenseNumber field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Driver entity, updatedBy field required")
	}
	return nil
}
