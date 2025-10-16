package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Maintenance struct {
	MaintenanceID   int32
	VehicleID       int32
	MaintenanceType string
	Description     string
	MaintenanceDate time.Time
	ServiceProvider string
	Cost            float64
	NextServiceDue  time.Time
	CreatedBy       int32
	UpdatedBy       int32
	DeletedBy       int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

// Constructor for creating a new Maintenance entity
func NewMaintenance(vehicleID int32, maintenanceType, description string, maintenanceDate time.Time, serviceProvider string, cost float64, nextServiceDue time.Time, createdBy int32) (*Maintenance, error) {
	entity := &Maintenance{
		VehicleID:       vehicleID,
		MaintenanceType: maintenanceType,
		Description:     description,
		MaintenanceDate: maintenanceDate,
		ServiceProvider: serviceProvider,
		Cost:            cost,
		NextServiceDue:  nextServiceDue,
		CreatedBy:       createdBy,
	}

	err := entity.ValidateNewMaintenance()
	if err != nil {
		log.Errorf("error validating new Maintenance entity %v", err)
		return &Maintenance{}, err
	}

	return entity, nil
}

// Validation for creating a new Maintenance
func (r *Maintenance) ValidateNewMaintenance() error {
	if r.MaintenanceDate.IsZero() {
		return errors.New("error validating Maintenance entity, maintenanceDate field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Maintenance entity, createdBy field required")
	}
	return nil
}

// Validation for updating an existing Maintenance
func (r *Maintenance) ValidateUpdateMaintenance() error {
	if r.MaintenanceID <= 0 {
		return errors.New("error validating Maintenance entity, maintenanceID field required")
	}
	if r.MaintenanceDate.IsZero() {
		return errors.New("error validating Maintenance entity, maintenanceDate field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Maintenance entity, updatedBy field required")
	}
	return nil
}
