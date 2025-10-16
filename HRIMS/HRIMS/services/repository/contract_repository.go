package repository

import (
	"context"
	"fmt"
	"os"
	"time"
	"training-backend/services/database"
	"training-backend/services/entity"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

// ContractConn initializes connection to DB
type ContractConn struct {
	conn *pgxpool.Pool
}

// NewContract connects to DB
func NewContract() *ContractConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &ContractConn{
		conn: conn,
	}
}

// contractSelectQuery returns base select query
func contractSelectQuery() string {
	return `SELECT
				id,
				application_id,
				employee_id,
				job_id,
				contract_type,
				start_date,
				end_date,
				signed_at,
				contract_status_id,
				resvd5,
				resvd4,
				resvd3,
				resvd2,
				resvd1,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM contract WHERE deleted_at IS NULL`
}

// Create inserts new contract
func (con *ContractConn) Create(e *entity.Contract) (int32, error) {
	var id int32
	query := `INSERT INTO contract
				(application_id, employee_id, job_id, contract_type, start_date, end_date, signed_at, contract_status_id, created_by, created_at)
			  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.ApplicationID,
		e.EmployeeID,
		e.JobID,
		e.ContractType,
		e.StartDate,
		e.EndDate,
		e.SignedAt,
		e.ContractStatusID,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// List lists all contracts
func (con *ContractConn) List() ([]*entity.Contract, error) {
	var (
		id, applicationID, employeeID, jobID, contractStatusID, createdBy, updatedBy, deletedBy pgtype.Int4
		contractType, resvd1, resvd2, resvd3, resvd4, resvd5                            pgtype.Text
		startDate, endDate, signedAt, createdAt, updatedAt, deletedAt                   pgtype.Timestamp
	)

	items := []*entity.Contract{}
	query := contractSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Contract{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&applicationID,
			&employeeID,
			&jobID,
			&contractType,
			&startDate,
			&endDate,
			&signedAt,
			&contractStatusID,
			&resvd5,
			&resvd4,
			&resvd3,
			&resvd2,
			&resvd1,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Contract{}, err
		}

		var endDatePtr, updatedAtPtr, deletedAtPtr *time.Time
		if endDate.Status == pgtype.Present {
			endDatePtr = &endDate.Time
		}
		if updatedAt.Status == pgtype.Present {
			updatedAtPtr = &updatedAt.Time
		}
		if deletedAt.Status == pgtype.Present {
			deletedAtPtr = &deletedAt.Time
		}

		item := &entity.Contract{
			ID:               id.Int,
			ApplicationID:    applicationID.Int,
			EmployeeID:       employeeID.Int,
			JobID:            jobID.Int,
			ContractType:     contractType.String,
			StartDate:        startDate.Time,
			EndDate:          endDatePtr,
			SignedAt:         signedAt.Time,
			ContractStatusID: contractStatusID.Int,
			Resvd1:           resvd1.String,
			Resvd2:           resvd2.String,
			Resvd3:           resvd3.String,
			Resvd4:           resvd4.String,
			Resvd5:           resvd5.String,
			CreatedBy:        createdBy.Int,
			UpdatedBy:        updatedBy.Int,
			DeletedBy:        deletedBy.Int,
			CreatedAt:        createdAt.Time,
			UpdatedAt:        updatedAtPtr,
			DeletedAt:        deletedAtPtr,
		}
		items = append(items, item)
	}

	return items, err
}

// Get retrieves single contract by ID
func (con *ContractConn) Get(id int32) (*entity.Contract, error) {
	var (
		applicationID, employeeID, jobID, contractStatusID, createdBy, updatedBy, deletedBy pgtype.Int4
		contractType, resvd1, resvd2, resvd3, resvd4, resvd5                            pgtype.Text
		startDate, endDate, signedAt, createdAt, updatedAt, deletedAt                   pgtype.Timestamp
	)
	var item *entity.Contract

	query := contractSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&applicationID,
		&employeeID,
		&jobID,
		&contractType,
		&startDate,
		&endDate,
		&signedAt,
		&contractStatusID,
		&resvd5,
		&resvd4,
		&resvd3,
		&resvd2,
		&resvd1,
		&createdBy,
		&updatedBy,
		&deletedBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return nil, err
	}

	var endDatePtr, updatedAtPtr, deletedAtPtr *time.Time
	if endDate.Status == pgtype.Present {
		endDatePtr = &endDate.Time
	}
	if updatedAt.Status == pgtype.Present {
		updatedAtPtr = &updatedAt.Time
	}
	if deletedAt.Status == pgtype.Present {
		deletedAtPtr = &deletedAt.Time
	}

	item = &entity.Contract{
		ID:               id,
		ApplicationID:    applicationID.Int,
		EmployeeID:       employeeID.Int,
		JobID:            jobID.Int,
		ContractType:     contractType.String,
		StartDate:        startDate.Time,
		EndDate:          endDatePtr,
		SignedAt:         signedAt.Time,
		ContractStatusID: contractStatusID.Int,
		Resvd1:           resvd1.String,
		Resvd2:           resvd2.String,
		Resvd3:           resvd3.String,
		Resvd4:           resvd4.String,
		Resvd5:           resvd5.String,
		CreatedBy:        createdBy.Int,
		UpdatedBy:        updatedBy.Int,
		DeletedBy:        deletedBy.Int,
		CreatedAt:        createdAt.Time,
		UpdatedAt:        updatedAtPtr,
		DeletedAt:        deletedAtPtr,
	}

	return item, nil
}

// Update updates a contract
func (con *ContractConn) Update(e *entity.Contract) (int32, error) {
	query := `UPDATE contract SET 
				application_id=$1,
				employee_id=$2,
				job_id=$3,
				contract_type=$4,
				start_date=$5,
				end_date=$6,
				contract_status_id=$7,
				updated_by=$8,
				updated_at=$9
			  WHERE id=$10`
	_, err := con.conn.Exec(context.Background(), query,
		e.ApplicationID,
		e.EmployeeID,
		e.JobID,
		e.ContractType,
		e.StartDate,
		e.EndDate,
		e.ContractStatusID,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	return e.ID, err
}

// SoftDelete marks contract as deleted
func (con *ContractConn) SoftDelete(id, deletedBy int32) error {
	query := `UPDATE contract SET deleted_by=$1, deleted_at=$2 WHERE id=$3`
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete permanently deletes contract
func (con *ContractConn) HardDelete(id int32) error {
	query := `DELETE FROM contract WHERE id=$1`
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
