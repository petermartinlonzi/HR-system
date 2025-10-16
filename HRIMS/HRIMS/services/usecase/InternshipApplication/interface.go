package internshipapplication

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.InternshipApplication, error)
	List() ([]*entity.InternshipApplication, error)
}

type Writer interface {
	Create(e *entity.InternshipApplication) (int32, error)
	Update(e *entity.InternshipApplication) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateInternshipApplication(studentID, departmentID, createdBy int32, resume string) (int32, error)
	ListInternshipApplications() ([]*entity.InternshipApplication, error)
	GetInternshipApplication(id int32) (*entity.InternshipApplication, error)
	UpdateInternshipApplication(e *entity.InternshipApplication) (int32, error)
	SoftDeleteInternshipApplication(id, deletedBy int32) error
	HardDeleteInternshipApplication(id int32) error
}
