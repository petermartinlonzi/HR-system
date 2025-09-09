package leavebalance

import (
	"database/sql"
	"training-backend/package/log"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
)

type Service struct {
	repo Repository
}

func NewService(db *sql.DB) UseCase {
	repo := repository.NewLeaveBalanceRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateLeaveBalance(userID int32, leaveType string, year, totalEntitled, createdBy int32) (int32, error) {
	var id int32
	e, err := entity.NewLeaveBalance(userID, leaveType, year, totalEntitled, createdBy)
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

func (s *Service) ListLeaveBalances() ([]*entity.LeaveBalance, error) {
	return s.repo.List()
}

func (s *Service) GetLeaveBalance(id int32) (*entity.LeaveBalance, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateLeaveBalance(e *entity.LeaveBalance) (int32, error) {
	err := e.ValidateUpdateLeaveBalance()
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

func (s *Service) SoftDeleteLeaveBalance(id, deletedBy int32) error {
	_, err := s.GetLeaveBalance(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteLeaveBalance(id int32) error {
	_, err := s.GetLeaveBalance(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
