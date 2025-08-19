package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type TransportRequest struct {
	RequestID    int32
	RequesterID  int32
	DriverID     int32
	VehicleID    int32
	Origin       string
	Destination  string
	Purpose      string
	RequestedDate time.Time
	DepartureTime time.Time
	ReturnTime    time.Time
	ApprovalStatus string

	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewTransportRequest(requesterID int32, origin, destination, approvalStatus string, requestedDate time.Time, createdBy int32) (*TransportRequest, error) {
	entity := &TransportRequest{
		RequesterID:    requesterID,
		Origin:         origin,
		Destination:    destination,
		ApprovalStatus: approvalStatus,
		RequestedDate:  requestedDate,
		CreatedBy:      createdBy,
	}

	err := entity.ValidateNewTransportRequest()
	if err != nil {
		log.Errorf("error validating new TransportRequest entity %v", err)
		return &TransportRequest{}, err
	}

	return entity, nil
}

func (r *TransportRequest) ValidateNewTransportRequest() error {
	if r.RequesterID <= 0 {
		return errors.New("requesterID is required")
	}
	if r.Origin == "" || r.Destination == "" {
		return errors.New("origin and destination are required")
	}
	if r.RequestedDate.IsZero() {
		return errors.New("requestedDate is required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("createdBy is required")
	}
	return nil
}

func (r *TransportRequest) ValidateUpdateTransportRequest() error {
	if r.RequestID <= 0 {
		return errors.New("requestID is required for update")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("updatedBy is required")
	}
	return nil
}
