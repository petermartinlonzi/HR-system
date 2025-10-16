package fuelrecord

import (
	"database/sql"
	"time"
	"training-backend/package/log"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
)

type Service struct {
	repo Repository
}

func NewService(db *sql.DB) UseCase {
	repo := repository.NewFuelRecordsRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateFuelRecord(vehicleID, createdBy int32, fuelingDate string, quantity float64, fuelType string) (int32, error) {
	var id int32
	date, _ := time.Parse(time.RFC3339, fuelingDate)
	e, err := entity.NewFuelRecord(vehicleID, createdBy, date, quantity, fuelType)
	if err != nil {
		log.Error(err)
		return id, err
	}
	id, err = s.repo.Create(e)
	if err != nil {
		log.Error(err)
		return id, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListFuelRecords() ([]*entity.FuelRecord, error) {
	return s.repo.List()
}

func (s *Service) GetFuelRecord(id int32) (*entity.FuelRecord, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateFuelRecord(e *entity.FuelRecord) (int32, error) {
	err := e.ValidateUpdateFuelRecord()
	if err != nil {
		log.Error(err)
		return e.FuelID, error_message.ErrCannotBeUpdated
	}
	_, err = s.repo.Update(e)
	if err != nil {
		log.Error(err)
		return e.FuelID, error_message.ErrNotFound
	}
	return e.FuelID, nil
}

func (s *Service) SoftDeleteFuelRecord(id, deletedBy int32) error {
	_, err := s.GetFuelRecord(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteFuelRecord(id int32) error {
	_, err := s.GetFuelRecord(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
