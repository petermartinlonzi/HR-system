package user

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.User, error)
	List() ([]*entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (int32, error)
	Update(e *entity.User) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateUser(firstName, email, passwordHash string, roleID, createdBy int32) (int32, error)
	ListUsers() ([]*entity.User, error)
	GetUser(id int32) (*entity.User, error)
	UpdateUser(e *entity.User) (int32, error)
	SoftDeleteUser(id, deletedBy int32) error
	HardDeleteUser(id int32) error
}
