package internshipapplication

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
	repo := repository.NewInternshipApplicationsRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateInternshipApplication(studentID, departmentID, createdBy int32, resume string) (int32, error) {
	var id int32
	e, err := entity.NewInternshipApplication(studentID, departmentID, createdBy, resume)
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

func (s *Service) ListInternshipApplications() ([]*entity.InternshipApplication, error) {
	return s.repo.List()
}

func (s *Service) GetInternshipApplication(id int32) (*entity.InternshipApplication, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateInternshipApplication(e *entity.InternshipApplication) (int32, error) {
	err := e.ValidateUpdateInternshipApplication()
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

func (s *Service) SoftDeleteInternshipApplication(id, deletedBy int32) error {
	_, err := s.GetInternshipApplication(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteInternshipApplication(id int32) error {
	_, err := s.GetInternshipApplication(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
