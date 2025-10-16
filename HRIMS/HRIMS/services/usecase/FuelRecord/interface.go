package fuelrecord

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.FuelRecord, error)
	List() ([]*entity.FuelRecord, error)
}

type Writer interface {
	Create(e *entity.FuelRecord) (int32, error)
	Update(e *entity.FuelRecord) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateFuelRecord(vehicleID, createdBy int32, fuelingDate string, quantity float64, fuelType string) (int32, error)
	ListFuelRecords() ([]*entity.FuelRecord, error)
	GetFuelRecord(id int32) (*entity.FuelRecord, error)
	UpdateFuelRecord(e *entity.FuelRecord) (int32, error)
	SoftDeleteFuelRecord(id, deletedBy int32) error
	HardDeleteFuelRecord(id int32) error
}
