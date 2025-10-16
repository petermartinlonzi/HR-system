package user

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
	repo := repository.NewUserRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateUser(firstName, email, passwordHash string, roleID, createdBy int32) (int32, error) {
	var id int32
	user, err := entity.NewUser(firstName, email, passwordHash, roleID, createdBy)
	if err != nil {
		log.Error(err)
		return id, err
	}
	id, err = s.repo.Create(user)
	if err != nil {
		log.Error(err)
		return id, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}

func (s *Service) GetUser(id int32) (*entity.User, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateUser(e *entity.User) (int32, error) {
	err := e.ValidateUpdateUser()
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

func (s *Service) SoftDeleteUser(id, deletedBy int32) error {
	_, err := s.GetUser(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteUser(id int32) error {
	_, err := s.GetUser(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
