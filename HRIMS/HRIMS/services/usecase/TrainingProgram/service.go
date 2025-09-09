package training_program

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


// NewService creates a new Service with the given *sql.DB.
func NewService(db *sql.DB) UseCase {
	repo := repository.NewTrainingProgramRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateTrainingProgram(title, description string, departmentID, createdBy int32) (int32, error) {
	var id int32
	program, err := entity.NewTrainingProgram(title, description, departmentID, createdBy)
	if err != nil {
		log.Error(err)
		return id, err
	}
	id, err = s.repo.Create(program)
	if err != nil {
		log.Error(err)
		return id, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListTrainingPrograms() ([]*entity.TrainingProgram, error) {
	return s.repo.List()
}

func (s *Service) GetTrainingProgram(id int32) (*entity.TrainingProgram, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateTrainingProgram(e *entity.TrainingProgram) (int32, error) {
	err := e.ValidateUpdateTrainingProgram()
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

func (s *Service) SoftDeleteTrainingProgram(id, deletedBy int32) error {
	_, err := s.GetTrainingProgram(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteTrainingProgram(id int32) error {
	_, err := s.GetTrainingProgram(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
