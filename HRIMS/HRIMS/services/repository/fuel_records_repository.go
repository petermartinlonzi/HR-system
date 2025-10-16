package repository

import (
	"database/sql"
	"errors"
	"time"
	"training-backend/services/entity"
	"training-backend/package/log"
)

// FuelRecordsRepository manages CRUD operations for fuel records
type FuelRecordsRepository struct {
	db *sql.DB
}

// NewFuelRecordsRepository initializes the repository
func NewFuelRecordsRepository(db *sql.DB) *FuelRecordsRepository {
	return &FuelRecordsRepository{db: db}
}

// Create inserts a new fuel record into the database
func (r *FuelRecordsRepository) Create(record *entity.FuelRecord) (*entity.FuelRecord, error) {
	query := `
		INSERT INTO fuel_records 
			(vehicle_id, fueling_date, fuel_type, quantity_liters, cost, odometer_reading, fueling_station, created_by, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING fuel_id, created_at;
	`

	err := r.db.QueryRow(query,
		record.VehicleID,
		record.FuelingDate,
		record.FuelType,
		record.QuantityLiters,
		record.Cost,
		record.OdometerReading,
		record.FuelingStation,
		record.CreatedBy,
		time.Now(),
	).Scan(&record.FuelID, &record.CreatedAt)

	if err != nil {
		log.Errorf("failed to create fuel record: %v", err)
		return nil, err
	}

	return record, nil
}

// GetByID fetches a single fuel record by ID
func (r *FuelRecordsRepository) GetByID(fuelID int32) (*entity.FuelRecord, error) {
	query := `
		SELECT fuel_id, vehicle_id, fueling_date, fuel_type, quantity_liters, cost, odometer_reading,
		       fueling_station, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM fuel_records
		WHERE fuel_id = $1 AND deleted_at IS NULL;
	`

	record := &entity.FuelRecord{}
	err := r.db.QueryRow(query, fuelID).Scan(
		&record.FuelID,
		&record.VehicleID,
		&record.FuelingDate,
		&record.FuelType,
		&record.QuantityLiters,
		&record.Cost,
		&record.OdometerReading,
		&record.FuelingStation,
		&record.CreatedBy,
		&record.UpdatedBy,
		&record.DeletedBy,
		&record.CreatedAt,
		&record.UpdatedAt,
		&record.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		log.Errorf("failed to get fuel record by id: %v", err)
		return nil, err
	}

	return record, nil
}

// GetAll retrieves all active fuel records
func (r *FuelRecordsRepository) GetAll() ([]*entity.FuelRecord, error) {
	query := `
		SELECT fuel_id, vehicle_id, fueling_date, fuel_type, quantity_liters, cost, odometer_reading,
		       fueling_station, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM fuel_records
		WHERE deleted_at IS NULL
		ORDER BY fueling_date DESC;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("failed to fetch fuel records: %v", err)
		return nil, err
	}
	defer rows.Close()

	var records []*entity.FuelRecord
	for rows.Next() {
		record := &entity.FuelRecord{}
		err := rows.Scan(
			&record.FuelID,
			&record.VehicleID,
			&record.FuelingDate,
			&record.FuelType,
			&record.QuantityLiters,
			&record.Cost,
			&record.OdometerReading,
			&record.FuelingStation,
			&record.CreatedBy,
			&record.UpdatedBy,
			&record.DeletedBy,
			&record.CreatedAt,
			&record.UpdatedAt,
			&record.DeletedAt,
		)
		if err != nil {
			log.Errorf("failed to scan fuel record: %v", err)
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// Update modifies an existing fuel record
func (r *FuelRecordsRepository) Update(record *entity.FuelRecord) error {
	if err := record.ValidateUpdateFuelRecord(); err != nil {
		return err
	}

	query := `
		UPDATE fuel_records
		SET fueling_date = $1, fuel_type = $2, quantity_liters = $3, cost = $4, odometer_reading = $5,
		    fueling_station = $6, updated_by = $7, updated_at = $8
		WHERE fuel_id = $9 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query,
		record.FuelingDate,
		record.FuelType,
		record.QuantityLiters,
		record.Cost,
		record.OdometerReading,
		record.FuelingStation,
		record.UpdatedBy,
		time.Now(),
		record.FuelID,
	)

	if err != nil {
		log.Errorf("failed to update fuel record: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("fuel record not found or already deleted")
	}

	return nil
}

// Delete performs a soft delete on a fuel record
func (r *FuelRecordsRepository) Delete(fuelID, deletedBy int32) error {
	query := `
		UPDATE fuel_records
		SET deleted_by = $1, deleted_at = $2
		WHERE fuel_id = $3 AND deleted_at IS NULL;
	`

	result, err := r.db.Exec(query, deletedBy, time.Now(), fuelID)
	if err != nil {
		log.Errorf("failed to delete fuel record: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("fuel record not found or already deleted")
	}

	return nil
}
