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

// ContractVersionConn initializes connection to DB
type ContractVersionConn struct {
	conn *pgxpool.Pool
}

// NewContractVersion connects to DB
func NewContractVersion() *ContractVersionConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &ContractVersionConn{
		conn: conn,
	}
}

// Create inserts new ContractVersion record to DB
func (con *ContractVersionConn) Create(e *entity.ContractVersion) (int32, error) {
	var id int32
	query := `INSERT INTO contract_versions
	(contract_id, version_number, salary, benefits, working_hours, probation_period, signed_by, signed_at, created_by, created_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.ContractID,
		e.VersionNumber,
		e.Salary,
		e.Benefits,
		e.WorkingHours,
		e.ProbationPeriod,
		e.SignedBy,
		e.SignedAt,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// contractVersionSelectQuery returns base SELECT query
func contractVersionSelectQuery() string {
	return `SELECT
				id,
				contract_id,
				version_number,
				salary,
				benefits,
				working_hours,
				probation_period,
				signed_by,
				signed_at,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM contract_versions WHERE deleted_at IS NULL`
}

// List lists all contract versions
func (con *ContractVersionConn) List() ([]*entity.ContractVersion, error) {
	var id, contractID, versionNumber, signedBy, createdBy, updatedBy, deletedBy pgtype.Int4
	var salary pgtype.Float8
	var benefits, workingHours, probationPeriod pgtype.Text
	var signedAt, createdAt, updatedAt, deletedAt pgtype.Timestamp

	var items []*entity.ContractVersion

	query := contractVersionSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.ContractVersion{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&contractID,
			&versionNumber,
			&salary,
			&benefits,
			&workingHours,
			&probationPeriod,
			&signedBy,
			&signedAt,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.ContractVersion{}, err
		}

		item := &entity.ContractVersion{
			ID:              id.Int,
			ContractID:      contractID.Int,
			VersionNumber:   versionNumber.Int,
			Salary:          salary.Float,
			Benefits:        benefits.String,
			WorkingHours:    workingHours.String,
			ProbationPeriod: probationPeriod.String,
			SignedBy:        signedBy.Int,
			SignedAt:        signedAt.Time.Local(),
			CreatedBy:       createdBy.Int,
			UpdatedBy:       updatedBy.Int,
			DeletedBy:       deletedBy.Int,
			CreatedAt:       createdAt.Time.Local(),
			UpdatedAt:       &updatedAt.Time,
			DeletedAt:       &deletedAt.Time,
		}
		items = append(items, item)
	}

	return items, err
}

// Get gets single contract version record by ID
func (con *ContractVersionConn) Get(id int32) (*entity.ContractVersion, error) {
	var contractID, versionNumber, signedBy, createdBy, updatedBy, deletedBy pgtype.Int4
	var salary pgtype.Float8
	var benefits, workingHours, probationPeriod pgtype.Text
	var signedAt, createdAt, updatedAt, deletedAt pgtype.Timestamp

	var item *entity.ContractVersion

	query := contractVersionSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&contractID,
		&versionNumber,
		&salary,
		&benefits,
		&workingHours,
		&probationPeriod,
		&signedBy,
		&signedAt,
		&createdBy,
		&updatedBy,
		&deletedBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		return item, err
	}

	item = &entity.ContractVersion{
		ID:              id,
		ContractID:      contractID.Int,
		VersionNumber:   versionNumber.Int,
		Salary:          salary.Float,
		Benefits:        benefits.String,
		WorkingHours:    workingHours.String,
		ProbationPeriod: probationPeriod.String,
		SignedBy:        signedBy.Int,
		SignedAt:        signedAt.Time.Local(),
		CreatedBy:       createdBy.Int,
		UpdatedBy:       updatedBy.Int,
		DeletedBy:       deletedBy.Int,
		CreatedAt:       createdAt.Time.Local(),
		UpdatedAt:       &updatedAt.Time,
		DeletedAt:       &deletedAt.Time,
	}
	return item, err
}

// Update updates single contract version record
func (con *ContractVersionConn) Update(e *entity.ContractVersion) (int32, error) {
	query := `UPDATE contract_versions SET
				contract_id = $1,
				version_number = $2,
				salary = $3,
				benefits = $4,
				working_hours = $5,
				probation_period = $6,
				signed_by = $7,
				signed_at = $8,
				updated_by = $9,
				updated_at = $10
			WHERE id = $11`
	_, err := con.conn.Exec(context.Background(), query,
		e.ContractID,
		e.VersionNumber,
		e.Salary,
		e.Benefits,
		e.WorkingHours,
		e.ProbationPeriod,
		e.SignedBy,
		e.SignedAt,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// SoftDelete softly delete contract version record
func (con *ContractVersionConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE contract_versions SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	return err
}

// HardDelete permanently delete contract version record
func (con *ContractVersionConn) HardDelete(id int32) error {
	query := "DELETE FROM contract_versions WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	return err
}
