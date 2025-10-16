package equipmentissue

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
	repo := repository.NewEquipmentIssuesRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateEquipmentIssue(equipmentID, issuedTo, createdBy int32, issueDate string) (int32, error) {
	var id int32
	date, _ := time.Parse(time.RFC3339, issueDate)
	e, err := entity.NewEquipmentIssue(equipmentID, issuedTo, createdBy, date)
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

func (s *Service) ListEquipmentIssues() ([]*entity.EquipmentIssue, error) {
	return s.repo.List()
}

func (s *Service) GetEquipmentIssue(id int32) (*entity.EquipmentIssue, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateEquipmentIssue(e *entity.EquipmentIssue) (int32, error) {
	err := e.ValidateUpdateEquipmentIssue()
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

func (s *Service) SoftDeleteEquipmentIssue(id, deletedBy int32) error {
	_, err := s.GetEquipmentIssue(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteEquipmentIssue(id int32) error {
	_, err := s.GetEquipmentIssue(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
