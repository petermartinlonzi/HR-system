package student_training_enrollment

import "training-backend/services/entity"

// Reader interface
type Reader interface {
	Get(id int32) (*entity.StudentTrainingEnrollment, error)
	List() ([]*entity.StudentTrainingEnrollment, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.StudentTrainingEnrollment) (int32, error)
	Update(e *entity.StudentTrainingEnrollment) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreateStudentTrainingEnrollment(studentID, trainingProgramID, createdBy int32, status string) (int32, error)
	ListStudentTrainingEnrollments() ([]*entity.StudentTrainingEnrollment, error)
	GetStudentTrainingEnrollment(id int32) (*entity.StudentTrainingEnrollment, error)
	UpdateStudentTrainingEnrollment(e *entity.StudentTrainingEnrollment) (int32, error)
	SoftDeleteStudentTrainingEnrollment(id, deletedBy int32) error
	HardDeleteStudentTrainingEnrollment(id int32) error
}
