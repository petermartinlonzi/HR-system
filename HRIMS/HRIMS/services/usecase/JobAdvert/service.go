package jobadvert

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
	repo := repository.NewJobAdvertsRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateJobAdvert(title, description string, departmentID, postedBy int32, deadline string, createdBy, updatedBy, deletedBy string) (int32, error) {
	var id int32
	deadlineTime, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		log.Error(err)
		return id, err
	}

	e, err := entity.NewJobAdvert(title, description, departmentID, postedBy, 0, deadlineTime)
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

func (s *Service) ListJobAdverts() ([]*entity.JobAdvert, error) {
	return s.repo.List()
}

func (s *Service) GetJobAdvert(id int32) (*entity.JobAdvert, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateJobAdvert(e *entity.JobAdvert) (int32, error) {
	err := e.ValidateUpdateJobAdvert()
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

func (s *Service) SoftDeleteJobAdvert(id int32, deletedBy string) error {
	_, err := s.GetJobAdvert(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteJobAdvert(id int32) error {
	_, err := s.GetJobAdvert(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
