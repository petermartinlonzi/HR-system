package transport_request


import (
	"database/sql"
	"training-backend/package/log"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(db *sql.DB) UseCase {
	repo := repository.NewTransportRequestRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateTransportRequest(requesterID int32, origin, destination, approvalStatus string, requestedDate int64, createdBy int32) (int32, error) {
	var id int32
	date := time.Unix(requestedDate, 0)
	request, err := entity.NewTransportRequest(requesterID, origin, destination, approvalStatus, date, createdBy)
	if err != nil {
		log.Error(err)
		return id, err
	}
	id, err = s.repo.Create(request)
	if err != nil {
		log.Error(err)
		return id, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListTransportRequests() ([]*entity.TransportRequest, error) {
	return s.repo.List()
}

func (s *Service) GetTransportRequest(id int32) (*entity.TransportRequest, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateTransportRequest(e *entity.TransportRequest) (int32, error) {
	err := e.ValidateUpdateTransportRequest()
	if err != nil {
		log.Error(err)
		return e.RequestID, error_message.ErrCannotBeUpdated
	}
	_, err = s.repo.Update(e)
	if err != nil {
		log.Error(err)
		return e.RequestID, error_message.ErrNotFound
	}
	return e.RequestID, nil
}

func (s *Service) SoftDeleteTransportRequest(id, deletedBy int32) error {
	_, err := s.GetTransportRequest(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteTransportRequest(id int32) error {
	_, err := s.GetTransportRequest(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
