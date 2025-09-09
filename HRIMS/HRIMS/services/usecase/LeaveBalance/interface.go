package leavebalance

import "training-backend/services/entity"

type Reader interface {
	Get(id int32) (*entity.LeaveBalance, error)
	List() ([]*entity.LeaveBalance, error)
}

type Writer interface {
	Create(e *entity.LeaveBalance) (int32, error)
	Update(e *entity.LeaveBalance) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateLeaveBalance(userID int32, leaveType string, year, totalEntitled, createdBy int32) (int32, error)
	ListLeaveBalances() ([]*entity.LeaveBalance, error)
	GetLeaveBalance(id int32) (*entity.LeaveBalance, error)
	UpdateLeaveBalance(e *entity.LeaveBalance) (int32, error)
	SoftDeleteLeaveBalance(id, deletedBy int32) error
	HardDeleteLeaveBalance(id int32) error
}
