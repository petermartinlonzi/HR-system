package team

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.Team, error)
	List() ([]*entity.Team, error)
}

type Writer interface {
	Create(e *entity.Team) (int32, error)
	Update(e *entity.Team) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTeam(teamName, description string, createdBy int32) (int32, error)
	ListTeams() ([]*entity.Team, error)
	GetTeam(id int32) (*entity.Team, error)
	UpdateTeam(e *entity.Team) (int32, error)
	SoftDeleteTeam(id, deletedBy int32) error
	HardDeleteTeam(id int32) error
}
