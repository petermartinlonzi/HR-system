package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type User struct {
	ID           int32
	FirstName    string
	MiddleName   string
	Surname      string
	Age          int32
	Email        string
	PhoneNumber  string
	Password     string
	PasswordHash string
	RoleID       int32
	IsActive     bool

	CreatedBy int32
	UpdatedBy int32
	DeletedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(firstName, email, passwordHash string, roleID, createdBy int32) (*User, error) {
	entity := &User{
		FirstName:    firstName,
		Email:        email,
		PasswordHash: passwordHash,
		RoleID:       roleID,
		IsActive:     true,
		CreatedBy:    createdBy,
	}

	err := entity.ValidateNewUser()
	if err != nil {
		log.Errorf("error validating new User entity %v", err)
		return &User{}, err
	}

	return entity, nil
}

func (r *User) ValidateNewUser() error {
	if r.FirstName == "" {
		return errors.New("firstName is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.PasswordHash == "" {
		return errors.New("passwordHash is required")
	}
	if r.RoleID <= 0 {
		return errors.New("roleID is required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("createdBy is required")
	}
	return nil
}

func (r *User) ValidateUpdateUser() error {
	if r.ID <= 0 {
		return errors.New("id is required for update")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("updatedBy is required")
	}
	return nil
}
