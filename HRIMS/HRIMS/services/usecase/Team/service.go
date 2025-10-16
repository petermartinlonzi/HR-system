package team

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
	repo := repository.NewTeamRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateTeam(teamName, description string, createdBy int32) (int32, error) {
	var id int32
	team, err := entity.NewTeam(teamName, description, createdBy)
	if err != nil {
		log.Error(err)
		return id, err
	}
	id, err = s.repo.Create(team)
	if err != nil {
		log.Error(err)
		return id, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListTeams() ([]*entity.Team, error) {
	return s.repo.List()
}

func (s *Service) GetTeam(id int32) (*entity.Team, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateTeam(e *entity.Team) (int32, error) {
	err := e.ValidateUpdateTeam()
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

func (s *Service) SoftDeleteTeam(id, deletedBy int32) error {
	_, err := s.GetTeam(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteTeam(id int32) error {
	_, err := s.GetTeam(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
