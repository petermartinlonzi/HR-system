package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type SportsEquipment struct {
	ID                  int32
	Name                string
	Quantity            int32
	Condition           string
	Location            string
	LastMaintenanceDate *time.Time
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
func NewSportsEquipment(name string, quantity int32, condition, location string, lastMaintenanceDate *time.Time, createdBy int32) (*SportsEquipment, error) {
	if condition == "" {
		condition = "Good" // default
	}

	entity := &SportsEquipment{
		Name:                name,
		Quantity:            quantity,
		Condition:           condition,
		Location:            location,
		LastMaintenanceDate: lastMaintenanceDate,
		CreatedBy:           createdBy,
		CreatedAt:           time.Now(),
	}

	if err := entity.ValidateNewSportsEquipment(); err != nil {
		log.Errorf("error validating new SportsEquipment entity %v", err)
		return nil, err
	}

	return entity, nil
}
func (e *SportsEquipment) ValidateNewSportsEquipment() error {
	if e.Name == "" {
		return errors.New("error validating SportsEquipment entity, name required")
	}
	if e.Quantity < 0 {
		return errors.New("error validating SportsEquipment entity, quantity cannot be negative")
	}
	if e.CreatedBy <= 0 {
		return errors.New("error validating SportsEquipment entity, createdBy required")
	}
	return nil
}
func (e *SportsEquipment) ValidateUpdateSportsEquipment() error {
	if e.ID <= 0 {
		return errors.New("error validating SportsEquipment entity, id required")
	}
	if e.Name == "" {
		return errors.New("error validating SportsEquipment entity, name required")
	}
	if e.UpdatedBy <= 0 {
		return errors.New("error validating SportsEquipment entity, updatedBy required")
	}
	if e.Quantity < 0 {
		return errors.New("error validating SportsEquipment entity, quantity cannot be negative")
	}
	return nil
}
func (e *SportsEquipment) MarkDeleted(userID int32) {
	e.DeletedBy = userID
	e.DeletedAt = time.Now()
}
