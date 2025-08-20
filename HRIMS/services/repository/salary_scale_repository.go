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

// SalaryScaleConn Initializes connection to DB
type SalaryScaleConn struct {
	conn *pgxpool.Pool
}

// NewSalaryScale Connects to DB
func NewSalaryScale() *SalaryScaleConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &SalaryScaleConn{
		conn: conn,
	}
}

// Create Inserts new record to DB
func (con *SalaryScaleConn) Create(e *entity.SalaryScale) (int32, error) {
	var id int32
	query := `INSERT INTO salary_scale 
				(salary_scale_name, position_id, minimum_salary, maximum_salary, currency_type, created_by, created_at) 
			  VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	err := con.conn.QueryRow(context.Background(), query,
		e.SalaryScaleName,
		e.PositionID,
		e.MinimumSalary,
		e.MaximumSalary,
		e.CurrencyType,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)
	return id, err
}

// CheckSalaryScale Checks if record exists in DB by salary_scale_name
func (con *SalaryScaleConn) CheckSalaryScale(name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM salary_scale WHERE salary_scale_name = $1)"
	err := con.conn.QueryRow(context.Background(), query, name).Scan(&exists)
	return exists, err
}

func salaryScaleSelectQuery() string {
	return `SELECT
				id,
				salary_scale_name,
				position_id,
				minimum_salary,
				maximum_salary,
				currency_type,
				created_by,
				updated_by,
				deleted_by,
				created_at,
				updated_at,
				deleted_at
			FROM salary_scale WHERE deleted_at IS NULL`
}

// List Lists all records
func (con *SalaryScaleConn) List() ([]*entity.SalaryScale, error) {
	var id pgtype.Int4
	var name pgtype.Text
	var positionID pgtype.Int4
	var minSalary pgtype.Float8
	var maxSalary pgtype.Float8
	var currency pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.SalaryScale

	query := salaryScaleSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.SalaryScale{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
			&positionID,
			&minSalary,
			&maxSalary,
			&currency,
			&createdBy,
			&updatedBy,
			&deletedBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.SalaryScale{}, err
		}

		item := &entity.SalaryScale{
			ID:              id.Int,
			SalaryScaleName: name.String,
			PositionID:      positionID.Int,
			MinimumSalary:   minSalary.Float,
			MaximumSalary:   maxSalary.Float,
			CurrencyType:    currency.String,
			CreatedBy:       createdBy.Int,
			UpdatedBy:       updatedBy.Int,
			DeletedBy:       deletedBy.Int,
			CreatedAt:       createdAt.Time.Local(),
			UpdatedAt:       updatedAt.Time.Local(),
			DeletedAt:       deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get Gets single record by ID field
func (con *SalaryScaleConn) Get(id int32) (*entity.SalaryScale, error) {
	var name pgtype.Text
	var positionID pgtype.Int4
	var minSalary pgtype.Float8
	var maxSalary pgtype.Float8
	var currency pgtype.Text
	var createdBy pgtype.Int4
	var updatedBy pgtype.Int4
	var deletedBy pgtype.Int4
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.SalaryScale

	query := salaryScaleSelectQuery() + ` AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&name,
		&positionID,
		&minSalary,
		&maxSalary,
		&currency,
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

	item = &entity.SalaryScale{
		ID:              id,
		SalaryScaleName: name.String,
		PositionID:      positionID.Int,
		MinimumSalary:   minSalary.Float,
		MaximumSalary:   maxSalary.Float,
		CurrencyType:    currency.String,
		CreatedBy:       createdBy.Int,
		UpdatedBy:       updatedBy.Int,
		DeletedBy:       deletedBy.Int,
		CreatedAt:       createdAt.Time.Local(),
		UpdatedAt:       updatedAt.Time.Local(),
		DeletedAt:       deletedAt.Time.Local(),
	}
	return item, err
}

// Update Updates single record by ID field
func (con *SalaryScaleConn) Update(e *entity.SalaryScale) (int32, error) {
	query := `UPDATE salary_scale SET 
				salary_scale_name = $1, 
				position_id = $2,
				minimum_salary = $3,
				maximum_salary = $4,
				currency_type = $5,
				updated_by = $6, 
				updated_at = $7 
			  WHERE id = $8`
	_, err := con.conn.Exec(context.Background(), query,
		e.SalaryScaleName,
		e.PositionID,
		e.MinimumSalary,
		e.MaximumSalary,
		e.CurrencyType,
		e.UpdatedBy,
		time.Now().Local(),
		e.ID,
	)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

// SoftDelete Softly delete single record by ID
func (con *SalaryScaleConn) SoftDelete(id, deletedBy int32) error {
	query := "UPDATE salary_scale SET deleted_by = $1, deleted_at = $2 WHERE id = $3"
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete Permanently delete single record by ID
func (con *SalaryScaleConn) HardDelete(id int32) error {
	query := "DELETE FROM salary_scale WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
