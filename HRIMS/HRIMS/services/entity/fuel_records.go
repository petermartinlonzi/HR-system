package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type FuelRecord struct {
	FuelID          int32
	VehicleID       int32
	FuelingDate     time.Time
	FuelType        string
	QuantityLiters  float64
	Cost            float64
	OdometerReading int32
	FuelingStation  string
	CreatedBy       int32
	UpdatedBy       int32
	DeletedBy       int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

func NewFuelRecord(vehicleID, createdBy int32, fuelingDate time.Time, quantity float64, fuelType string) (*FuelRecord, error) {
	entity := &FuelRecord{
		VehicleID:      vehicleID,
		FuelingDate:    fuelingDate,
		QuantityLiters: quantity,
		FuelType:       fuelType,
		CreatedBy:      createdBy,
	}

	err := entity.ValidateNewFuelRecord()
	if err != nil {
		log.Errorf("error validating new FuelRecord entity %v", err)
		return &FuelRecord{}, err
	}

	return entity, nil
}

func (r *FuelRecord) ValidateNewFuelRecord() error {
	if r.VehicleID <= 0 {
		return errors.New("error validating FuelRecord entity, vehicleID required")
	}
	if r.FuelingDate.IsZero() {
		return errors.New("error validating FuelRecord entity, fuelingDate required")
	}
	if r.QuantityLiters <= 0 {
		return errors.New("error validating FuelRecord entity, quantityLiters must be greater than 0")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating FuelRecord entity, createdBy required")
	}
	return nil
}

func (r *FuelRecord) ValidateUpdateFuelRecord() error {
	if r.FuelID <= 0 {
		return errors.New("error validating FuelRecord entity, fuelID required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating FuelRecord entity, updatedBy required")
	}
	return nil
}
