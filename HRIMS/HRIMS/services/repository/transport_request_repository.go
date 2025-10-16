package repository

import (
	"database/sql"
	"training-backend/services/entity"
	"training-backend/package/log"
)

type TransportRequestRepository struct {
	db *sql.DB
}

func NewTransportRequestRepository(db *sql.DB) *TransportRequestRepository {
	return &TransportRequestRepository{db: db}
}

func (r *TransportRequestRepository) Create(req *entity.TransportRequest) error {
	query := `
		INSERT INTO transport_requests
		(requester_id, driver_id, vehicle_id, origin, destination, purpose, requested_date, departure_time, return_time, approval_status, created_by, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,NOW())
		RETURNING request_id`
	return r.db.QueryRow(query, req.RequesterID, req.DriverID, req.VehicleID, req.Origin, req.Destination, req.Purpose, req.RequestedDate, req.DepartureTime, req.ReturnTime, req.ApprovalStatus, req.CreatedBy).Scan(&req.RequestID)
}

func (r *TransportRequestRepository) List() ([]*entity.TransportRequest, error) {
	query := `
		SELECT request_id, requester_id, driver_id, vehicle_id, origin, destination, purpose, requested_date, departure_time, return_time, approval_status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM transport_requests
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("error fetching transport requests: %v", err)
		return nil, err
	}
	defer rows.Close()

	var requests []*entity.TransportRequest
	for rows.Next() {
		tr := &entity.TransportRequest{}
		err := rows.Scan(&tr.RequestID, &tr.RequesterID, &tr.DriverID, &tr.VehicleID, &tr.Origin, &tr.Destination, &tr.Purpose, &tr.RequestedDate, &tr.DepartureTime, &tr.ReturnTime, &tr.ApprovalStatus, &tr.CreatedBy, &tr.UpdatedBy, &tr.DeletedBy, &tr.CreatedAt, &tr.UpdatedAt, &tr.DeletedAt)
		if err != nil {
			return nil, err
		}
		requests = append(requests, tr)
	}
	return requests, nil
}

func (r *TransportRequestRepository) Get(id int32) (*entity.TransportRequest, error) {
	query := `
		SELECT request_id, requester_id, driver_id, vehicle_id, origin, destination, purpose, requested_date, departure_time, return_time, approval_status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM transport_requests
		WHERE request_id = $1 AND deleted_at IS NULL`
	tr := &entity.TransportRequest{}
	err := r.db.QueryRow(query, id).Scan(&tr.RequestID, &tr.RequesterID, &tr.DriverID, &tr.VehicleID, &tr.Origin, &tr.Destination, &tr.Purpose, &tr.RequestedDate, &tr.DepartureTime, &tr.ReturnTime, &tr.ApprovalStatus, &tr.CreatedBy, &tr.UpdatedBy, &tr.DeletedBy, &tr.CreatedAt, &tr.UpdatedAt, &tr.DeletedAt)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (r *TransportRequestRepository) Update(req *entity.TransportRequest) error {
	query := `
		UPDATE transport_requests
		SET driver_id=$1, vehicle_id=$2, purpose=$3, approval_status=$4, updated_by=$5, updated_at=NOW()
		WHERE request_id=$6 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, req.DriverID, req.VehicleID, req.Purpose, req.ApprovalStatus, req.UpdatedBy, req.RequestID)
	return err
}

func (r *TransportRequestRepository) SoftDelete(id int32, deletedBy int32) error {
	query := `
		UPDATE transport_requests
		SET deleted_by=$1, deleted_at=NOW()
		WHERE request_id=$2 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, deletedBy, id)
	return err
}

func (r *TransportRequestRepository) HardDelete(id int32) error {
	_, err := r.db.Exec("DELETE FROM transport_requests WHERE request_id=$1", id)
	return err
}
