package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type SportsEvent struct {
	ID          int32
	EventName   string
	Description string
	Location    string
	EventDate   time.Time
	StartTime   *time.Time
	EndTime     *time.Time
	OrganizedBy int32
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
func NewSportsEvent(eventName, description, location string, eventDate time.Time, startTime, endTime *time.Time, organizedBy, createdBy int32) (*SportsEvent, error) {
	entity := &SportsEvent{
		EventName:   eventName,
		Description: description,
		Location:    location,
		EventDate:   eventDate,
		StartTime:   startTime,
		EndTime:     endTime,
		OrganizedBy: organizedBy,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
	}

	if err := entity.ValidateNewSportsEvent(); err != nil {
		log.Errorf("error validating new SportsEvent entity %v", err)
		return nil, err
	}

	return entity, nil
}
func (e *SportsEvent) ValidateNewSportsEvent() error {
	if e.EventName == "" {
		return errors.New("error validating SportsEvent entity, event_name required")
	}
	if e.EventDate.IsZero() {
		return errors.New("error validating SportsEvent entity, event_date required")
	}
	if e.CreatedBy <= 0 {
		return errors.New("error validating SportsEvent entity, createdBy required")
	}
	return nil
}
func (e *SportsEvent) ValidateUpdateSportsEvent() error {
	if e.ID <= 0 {
		return errors.New("error validating SportsEvent entity, id required")
	}
	if e.EventName == "" {
		return errors.New("error validating SportsEvent entity, event_name required")
	}
	if e.UpdatedBy <= 0 {
		return errors.New("error validating SportsEvent entity, updatedBy required")
	}
	return nil
}
func (e *SportsEvent) MarkDeleted(userID int32) {
	e.DeletedBy = userID
	e.DeletedAt = time.Now()
}
