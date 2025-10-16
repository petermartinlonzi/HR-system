package equipmentissue

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.EquipmentIssue, error)
	List() ([]*entity.EquipmentIssue, error)
}

type Writer interface {
	Create(e *entity.EquipmentIssue) (int32, error)
	Update(e *entity.EquipmentIssue) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateEquipmentIssue(equipmentID, issuedTo, createdBy int32, issueDate string) (int32, error)
	ListEquipmentIssues() ([]*entity.EquipmentIssue, error)
	GetEquipmentIssue(id int32) (*entity.EquipmentIssue, error)
	UpdateEquipmentIssue(e *entity.EquipmentIssue) (int32, error)
	SoftDeleteEquipmentIssue(id, deletedBy int32) error
	HardDeleteEquipmentIssue(id int32) error
}
