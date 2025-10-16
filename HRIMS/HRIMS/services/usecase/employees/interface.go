package employee

import "training-backend/services/entity"

// If the import path is incorrect or the Employees struct does not exist, define it here for now:
// Remove the below definition if you fix the import path and Employees struct exists in the entity package.
type Employees struct {
	ID        int32
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Age       int32
	CreatedBy int32
}

type Reader interface {
	Get(id int32) (*entity.Employees, error)
	List() ([]*entity.Employees, error)
}

type Writer interface {
	Create(e *entity.Employees) (int32, error)
	Update(e *entity.Employees) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateEmployees(firstName, lastName, email, phone string, age int32, createdBy int32) (int32, error)
	ListEmployeess() ([]*entity.Employees, error)
	GetEmployees(id int32) (*entity.Employees, error)
	UpdateEmployees(e *entity.Employees) (int32, error)
	SoftDeleteEmployees(id, deletedBy int32) error
	HardDeleteEmployees(id int32) error
}
