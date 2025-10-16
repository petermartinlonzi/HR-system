package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type AuditLog struct {
	ID        int32
	UserID    int32
	Action    string
	TableName string
	RecordID  int32
	Details   string
	CreatedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Constructor
func NewAuditLog(userID, recordID, createdBy int32, action, tableName, details string) (*AuditLog, error) {
	entity := &AuditLog{
		UserID:    userID,
		Action:    action,
		TableName: tableName,
		RecordID:  recordID,
		Details:   details,
		CreatedBy: createdBy,
	}

	err := entity.ValidateNewAuditLog()
	if err != nil {
		log.Errorf("error validating new AuditLog entity %v", err)
		return &AuditLog{}, err
	}

	return entity, nil
}

// Validation wakati wa ku-create audit log
func (r *AuditLog) ValidateNewAuditLog() error {
	if r.UserID <= 0 {
		return errors.New("error validating AuditLog entity, userID field required")
	}
	if r.Action == "" {
		return errors.New("error validating AuditLog entity, action field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating AuditLog entity, createdBy field required")
	}
	return nil
}

// Validation wakati wa ku-update audit log
func (r *AuditLog) ValidateUpdateAuditLog() error {
	if r.ID <= 0 {
		return errors.New("error validating AuditLog entity, id field required")
	}
	if r.Action == "" {
		return errors.New("error validating AuditLog entity, action field required")
	}
	return nil
}
