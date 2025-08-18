package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type SportsParticipant struct {
	ID              int32
	EventID         int32
	EmployeeID      int32
	TeamName        string
	PerformanceNotes string
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
func NewSportsParticipant(eventID, employeeID, createdBy int32, teamName, performanceNotes string) (*SportsParticipant, error) {
	entity := &SportsParticipant{
		EventID:          eventID,
		EmployeeID:       employeeID,
		TeamName:         teamName,
		PerformanceNotes: performanceNotes,
		CreatedBy:        createdBy,
		CreatedAt:        time.Now(),
	}

	if err := entity.ValidateNewSportsParticipant(); err != nil {
		log.Errorf("error validating new SportsParticipant entity %v", err)
		return nil, err
	}

	return entity, nil
}
func (p *SportsParticipant) ValidateNewSportsParticipant() error {
	if p.EventID <= 0 {
		return errors.New("error validating SportsParticipant entity, event_id required")
	}
	if p.EmployeeID <= 0 {
		return errors.New("error validating SportsParticipant entity, employee_id required")
	}
	if p.CreatedBy <= 0 {
		return errors.New("error validating SportsParticipant entity, createdBy required")
	}
	return nil
}
func (p *SportsParticipant) ValidateUpdateSportsParticipant() error {
	if p.ID <= 0 {
		return errors.New("error validating SportsParticipant entity, id required")
	}
	if p.UpdatedBy <= 0 {
		return errors.New("error validating SportsParticipant entity, updatedBy required")
	}
	return nil
}
func (p *SportsParticipant) MarkDeleted(userID int32) {
	p.DeletedBy = userID
	p.DeletedAt = time.Now()
}