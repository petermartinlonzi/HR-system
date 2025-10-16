package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Job_advert struct {
	ID          int32
	Title       string
	Description string
	DepartmentID int32
	PostedBy    int32
	Deadline    time.Time
    created_by string
    updated_by string
    deleted_by string
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
}

func NewJob_advert(title, description string, departmentID, postedBy int32, deadline time.Time, createdBy, updatedBy, deletedBy string, createdAt, updatedAt, deletedAt time.Time) (*Job_advert, error) {
	entity := &Job_advert{
		Title:       title,
		Description: description,
		DepartmentID: departmentID,
		PostedBy:    postedBy,
		Deadline:    deadline,
		created_by:  createdBy,
		updated_by:  updatedBy,
		deleted_by:  deletedBy,
		created_at:  createdAt,
		updated_at:  updatedAt,
		deleted_at:  deletedAt,
	}

	err := entity.ValidateNewJob_advert()
	if err != nil {
		log.Errorf("error validating new Job_advert entity %v", err)
		return &Job_advert{}, err
	}

	return entity, nil
}


func (r *Job_advert) ValidateNewJob_advert() error {
	if r.Title == "" {
		return errors.New("error validating Job_advert entity, title field required")
	}
	if r.created_by <= "" {
		return errors.New("error validating Job_advert entity, createdBy field required")
	}
	return nil
}

func (r *Job_advert) ValidateUpdateJob_advert() error {
	if r.ID <= 0 {
		return errors.New("error validating Job_advert entity, id field required")
	}
	if r.Title == "" {
		return errors.New("error validating Job_advert entity, title field required")
	}
	if r.updated_by <= "" {
		return errors.New("error validating Job_advert entity, updatedBy field required")
	}
	return nil
}
