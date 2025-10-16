package jobadvert

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.JobAdvert, error)
	List() ([]*entity.JobAdvert, error)
}

type Writer interface {
	Create(e *entity.JobAdvert) (int32, error)
	Update(e *entity.JobAdvert) (int32, error)
	SoftDelete(id int32, deletedBy string) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateJobAdvert(title, description string, departmentID, postedBy int32, deadline string, createdBy, updatedBy, deletedBy string) (int32, error)
	ListJobAdverts() ([]*entity.JobAdvert, error)
	GetJobAdvert(id int32) (*entity.JobAdvert, error)
	UpdateJobAdvert(e *entity.JobAdvert) (int32, error)
	SoftDeleteJobAdvert(id int32, deletedBy string) error
	HardDeleteJobAdvert(id int32) error
}
