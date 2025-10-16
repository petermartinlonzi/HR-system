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

// BudgetConn initializes DB connection for budget
type BudgetConn struct {
	conn *pgxpool.Pool
}

// NewBudget connects to DB
func NewBudget() *BudgetConn {
	conn, err := database.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &BudgetConn{
		conn: conn,
	}
}

// Create inserts a new budget record
func (con *BudgetConn) Create(e *entity.Budget) (int32, error) {
	var id int32
	query := `INSERT INTO budget
		(request_id, salary_scale_id, number_of_officers, total_budget, description, created_by, created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING budget_id`

	err := con.conn.QueryRow(context.Background(), query,
		e.RequestID,
		e.SalaryScaleID,
		e.NumberOfOfficers,
		e.TotalBudget,
		e.Description,
		e.CreatedBy,
		time.Now().Local(),
	).Scan(&id)

	return id, err
}

// budgetSelectQuery builds common SELECT query
func budgetSelectQuery() string {
	return `SELECT
				budget_id,
				request_id,
				salary_scale_id,
				number_of_officers,
				total_budget,
				description,
				created_by,
				created_at,
				updated_at,
				deleted_at
			FROM budget WHERE deleted_at IS NULL`
}

// List returns all budget records
func (con *BudgetConn) List() ([]*entity.Budget, error) {
	var budgetID pgtype.Int4
	var requestID pgtype.Int4
	var salaryScaleID pgtype.Int4
	var numberOfOfficers pgtype.Int4
	var totalBudget pgtype.Numeric
	var description pgtype.Text
	var createdBy pgtype.Text
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var items []*entity.Budget

	query := budgetSelectQuery()
	rows, err := con.conn.Query(context.Background(), query)
	if err != nil {
		return []*entity.Budget{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&budgetID,
			&requestID,
			&salaryScaleID,
			&numberOfOfficers,
			&totalBudget,
			&description,
			&createdBy,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return []*entity.Budget{}, err
		}

		// Convert pgtype.Numeric to float64
		var total float64
		if err := totalBudget.AssignTo(&total); err != nil {
			total = 0
		}

		item := &entity.Budget{
			BudgetID:        budgetID.Int,
			RequestID:       requestID.Int,
			SalaryScaleID:   salaryScaleID.Int,
			NumberOfOfficers: numberOfOfficers.Int,
			TotalBudget:     total,
			Description:     description.String,
			CreatedBy:       createdBy.String,
			CreatedAt:       createdAt.Time.Local(),
			UpdatedAt:       updatedAt.Time.Local(),
			DeletedAt:       deletedAt.Time.Local(),
		}
		items = append(items, item)
	}

	return items, err
}

// Get returns single budget by ID
func (con *BudgetConn) Get(id int32) (*entity.Budget, error) {
	var requestID pgtype.Int4
	var salaryScaleID pgtype.Int4
	var numberOfOfficers pgtype.Int4
	var totalBudget pgtype.Numeric
	var description pgtype.Text
	var createdBy pgtype.Text
	var createdAt pgtype.Timestamp
	var updatedAt pgtype.Timestamp
	var deletedAt pgtype.Timestamp

	var item *entity.Budget

	query := budgetSelectQuery() + ` AND budget_id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(
		&id,
		&requestID,
		&salaryScaleID,
		&numberOfOfficers,
		&totalBudget,
		&description,
		&createdBy,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return item, err
	}

	var total float64
	if err := totalBudget.AssignTo(&total); err != nil {
		total = 0
	}

	item = &entity.Budget{
		BudgetID:        id,
		RequestID:       requestID.Int,
		SalaryScaleID:   salaryScaleID.Int,
		NumberOfOfficers: numberOfOfficers.Int,
		TotalBudget:     total,
		Description:     description.String,
		CreatedBy:       createdBy.String,
		CreatedAt:       createdAt.Time.Local(),
		UpdatedAt:       updatedAt.Time.Local(),
		DeletedAt:       deletedAt.Time.Local(),
	}
	return item, err
}

// Update updates a budget record
func (con *BudgetConn) Update(e *entity.Budget) (int32, error) {
	query := `UPDATE budget SET 
				request_id = $1,
				salary_scale_id = $2,
				number_of_officers = $3,
				total_budget = $4,
				description = $5,
				updated_at = $6
			  WHERE budget_id = $7`
	_, err := con.conn.Exec(context.Background(), query,
		e.RequestID,
		e.SalaryScaleID,
		e.NumberOfOfficers,
		e.TotalBudget,
		e.Description,
		time.Now().Local(),
		e.BudgetID,
	)
	if err != nil {
		return e.BudgetID, err
	}
	return e.BudgetID, err
}

// SoftDelete marks a budget as deleted
func (con *BudgetConn) SoftDelete(id int32) error {
	query := "UPDATE budget SET deleted_at = $1 WHERE budget_id = $2"
	_, err := con.conn.Exec(context.Background(), query, time.Now().Local(), id)
	if err != nil {
		return err
	}
	return err
}

// HardDelete permanently deletes a budget record
func (con *BudgetConn) HardDelete(id int32) error {
	query := "DELETE FROM budget WHERE budget_id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return err
}
