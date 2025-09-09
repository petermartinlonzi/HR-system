package student_training_enrollment

import (
	"database/sql"
	_ "github.com/lib/pq"
	"training-backend/package/log"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
)

type Service struct {
	repo Repository
}


func NewService() UseCase {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error(err)
		return nil
	}
	repo := repository.NewStudentTrainingEnrollmentRepository(db)
	return &Service{repo: repo}
}

func (s *Service) CreateStudentTrainingEnrollment(studentID, trainingProgramID, createdBy int32, status string) (int32, error) {
	var id int32
	enrollment, err := entity.NewStudentTrainingEnrollment(studentID, trainingProgramID, createdBy, status)
	if err != nil {
		log.Error(err)
		return id, err
	}
	id, err = s.repo.Create(enrollment)
	if err != nil {
		log.Error(err)
		return id, error_message.ErrCannotBeCreated
	}
	return id, nil
}

func (s *Service) ListStudentTrainingEnrollments() ([]*entity.StudentTrainingEnrollment, error) {
	return s.repo.List()
}

func (s *Service) GetStudentTrainingEnrollment(id int32) (*entity.StudentTrainingEnrollment, error) {
	return s.repo.Get(id)
}

func (s *Service) UpdateStudentTrainingEnrollment(e *entity.StudentTrainingEnrollment) (int32, error) {
	err := e.ValidateUpdateStudentTrainingEnrollment()
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

func (s *Service) SoftDeleteStudentTrainingEnrollment(id, deletedBy int32) error {
	_, err := s.GetStudentTrainingEnrollment(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.SoftDelete(id, deletedBy)
}

func (s *Service) HardDeleteStudentTrainingEnrollment(id int32) error {
	_, err := s.GetStudentTrainingEnrollment(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	return s.repo.HardDelete(id)
}
