package transport_request

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.TransportRequest, error)
	List() ([]*entity.TransportRequest, error)
}

type Writer interface {
	Create(e *entity.TransportRequest) (int32, error)
	Update(e *entity.TransportRequest) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTransportRequest(requesterID int32, origin, destination, approvalStatus string, requestedDate int64, createdBy int32) (int32, error)
	ListTransportRequests() ([]*entity.TransportRequest, error)
	GetTransportRequest(id int32) (*entity.TransportRequest, error)
	UpdateTransportRequest(e *entity.TransportRequest) (int32, error)
	SoftDeleteTransportRequest(id, deletedBy int32) error
	HardDeleteTransportRequest(id int32) error
}
