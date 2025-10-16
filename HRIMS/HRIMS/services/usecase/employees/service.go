package employee

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
	repo := repository.NewEmployees()
	return &Service{repo: repo}
}

// Implements UseCase interface
func (s *Service) CreateEmployees(firstName, lastName, email, phone string, age int32, createdBy int32) (int32, error) {
	// Adjust parameters as needed to match entity.NewEmployees
	emp, err := entity.NewEmployees(firstName, lastName, email, 0, "", "", time.Now(), "", "", "")
	if err != nil {
		log.Error(err)
		return 0, err
	}
	id, err := s.repo.Create(emp)
	if err != nil {
		log.Error(err)
		return 0, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListEmployeess() ([]*entity.Employees, error) {
	return s.repo.List()
}

func (s *Service) GetEmployees(id int32) (*entity.Employees, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateEmployees(e *entity.Employees) (int32, error) {
	return s.repo.Update(e)
}

func (s *Service) SoftDeleteEmployees(id, deletedBy int32) error {
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteEmployees(id int32) error {
	return s.repo.HardDelete(id)
}

func (s *Service) ListEmployees() ([]*entity.Employees, error) {
	return s.repo.List()
}

func (s *Service) GetEmployee(id int32) (*entity.Employees, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateEmployee(e *entity.Employees) (int32, error) {
	err := e.ValidateUpdateEmployees()
	if err != nil {
		log.Error(err)
		return e.ID, error_message.ErrCannotBeUpdated
	}
	_, err = s.repo.Update(e)
	if err != nil {
		log.Error(err)
		return e.ID, error_message.ErrNotFound
	}
	return e.ID, nil
}

func (s *Service) SoftDeleteEmployee(id, deletedBy int32) error {
	_, err := s.GetEmployee(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteEmployee(id int32) error {
	_, err := s.GetEmployee(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
