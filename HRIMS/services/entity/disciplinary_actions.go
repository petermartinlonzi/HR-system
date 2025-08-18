package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Disciplinary_actions struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewDisciplinary_actions(name, description string, createdBy int32) (*Disciplinary_actions, error) {
	entity := &Disciplinary_actions{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewDisciplinary_actions()
	if err != nil {
		log.Errorf("error validating new Disciplinary_actions entity %v", err)
		return &Disciplinary_actions{}, err
	}

	return entity, nil
}

func (r *Disciplinary_actions) ValidateNewDisciplinary_actions() error {
	if r.Name == "" {
		return errors.New("error validating Disciplinary_actions entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Disciplinary_actions entity, createdBy field required")
	}
	return nil
}

func (r *Disciplinary_actions) ValidateUpdateDisciplinary_actions() error {
	if r.ID <= 0 {
		return errors.New("error validating Disciplinary_actions entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Disciplinary_actions entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Disciplinary_actions entity, updatedBy field required")
	}
	return nil
}
