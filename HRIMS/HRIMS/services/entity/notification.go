package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Notification struct {
	ID        int32
	UserID    int32
	Message   string
	IsRead    bool
	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Constructor for creating a new Notification entity
func NewNotification(userID int32, message string, createdBy int32) (*Notification, error) {
	entity := &Notification{
		UserID:    userID,
		Message:   message,
		IsRead:    false, // default false
		CreatedBy: createdBy,
	}

	err := entity.ValidateNewNotification()
	if err != nil {
		log.Errorf("error validating new Notification entity %v", err)
		return &Notification{}, err
	}

	return entity, nil
}

// Validation for creating a new Notification
func (r *Notification) ValidateNewNotification() error {
	if r.Message == "" {
		return errors.New("error validating Notification entity, message field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Notification entity, createdBy field required")
	}
	return nil
}

// Validation for updating an existing Notification
func (r *Notification) ValidateUpdateNotification() error {
	if r.ID <= 0 {
		return errors.New("error validating Notification entity, id field required")
	}
	if r.Message == "" {
		return errors.New("error validating Notification entity, message field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Notification entity, updatedBy field required")
	}
	return nil
}
