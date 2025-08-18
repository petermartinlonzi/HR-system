package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Request struct {
	ID           int32
	OfficerID    int32
	Title        string
	Content      string
	Status       string
	SubmittedAt  time.Time
	PositionID   *int32
	DepartmentID *int32
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewRequest(title, content string, officerID, createdBy int32, positionID, departmentID *int32) (*Request, error) {
	entity := &Request{
		OfficerID:    officerID,
		Title:        title,
		Content:      content,
		Status:       "Pending",
		SubmittedAt:  time.Now(),
		PositionID:   positionID,
		DepartmentID: departmentID,
		CreatedBy:    createdBy,
		CreatedAt:    time.Now(),
	}

	if err := entity.ValidateNewRequest(); err != nil {
		log.Errorf("error validating new Request entity %v", err)
		return nil, err
	}

	return entity, nil
}

func (r *Request) ValidateNewRequest() error {
	if r.Title == "" {
		return errors.New("error validating Request entity, title field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Request entity, createdBy field required")
	}
	if r.OfficerID <= 0 {
		return errors.New("error validating Request entity, officerID field required")
	}
	return nil
}

func (r *Request) ValidateUpdateRequest() error {
	if r.ID <= 0 {
		return errors.New("error validating Request entity, id field required")
	}
	if r.Title == "" {
		return errors.New("error validating Request entity, title field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Request entity, updatedBy field required")
	}
	return nil
}

func (r *Request) MarkDeleted(userID int32) {
	r.DeletedBy = userID
	r.DeletedAt = time.Now()
}
