package training_program

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.TrainingProgram, error)
	List() ([]*entity.TrainingProgram, error)
}

type Writer interface {
	Create(e *entity.TrainingProgram) (int32, error)
	Update(e *entity.TrainingProgram) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTrainingProgram(title, description string, departmentID, createdBy int32) (int32, error)
	ListTrainingPrograms() ([]*entity.TrainingProgram, error)
	GetTrainingProgram(id int32) (*entity.TrainingProgram, error)
	UpdateTrainingProgram(e *entity.TrainingProgram) (int32, error)
	SoftDeleteTrainingProgram(id, deletedBy int32) error
	HardDeleteTrainingProgram(id int32) error
}
