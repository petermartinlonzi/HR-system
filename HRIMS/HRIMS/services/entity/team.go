package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Team struct {
	ID          int32
	TeamName    string
	Description string
	CreatedBy   int32

	CreatedByUser int32
	UpdatedBy     int32
	DeletedBy     int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func NewTeam(teamName, description string, createdBy int32) (*Team, error) {
	entity := &Team{
		TeamName:    teamName,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewTeam()
	if err != nil {
		log.Errorf("error validating new Team entity %v", err)
		return &Team{}, err
	}

	return entity, nil
}

func (r *Team) ValidateNewTeam() error {
	if r.TeamName == "" {
		return errors.New("teamName is required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("createdBy is required")
	}
	return nil
}

func (r *Team) ValidateUpdateTeam() error {
	if r.ID <= 0 {
		return errors.New("id is required for update")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("updatedBy is required")
	}
	return nil
}
