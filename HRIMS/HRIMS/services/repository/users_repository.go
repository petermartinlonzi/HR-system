package repository

import (
	"database/sql"
	"training-backend/services/entity"
	"training-backend/package/log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	query := `
		INSERT INTO users
		(first_name, middle_name, surname, age, email, phone_number, password_hash, role_id, is_active, created_by, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())
		RETURNING id`
	return r.db.QueryRow(query, user.FirstName, user.MiddleName, user.Surname, user.Age, user.Email, user.PhoneNumber, user.PasswordHash, user.RoleID, user.IsActive, user.CreatedBy).Scan(&user.ID)
}

func (r *UserRepository) List() ([]*entity.User, error) {
	query := `
		SELECT id, first_name, middle_name, surname, age, email, phone_number, password_hash, role_id, is_active, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM users
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("error fetching users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		u := &entity.User{}
		err := rows.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.Surname, &u.Age, &u.Email, &u.PhoneNumber, &u.PasswordHash, &u.RoleID, &u.IsActive, &u.CreatedBy, &u.UpdatedBy, &u.DeletedBy, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) Get(id int32) (*entity.User, error) {
	query := `
		SELECT id, first_name, middle_name, surname, age, email, phone_number, password_hash, role_id, is_active, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL`
	u := &entity.User{}
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.Surname, &u.Age, &u.Email, &u.PhoneNumber, &u.PasswordHash, &u.RoleID, &u.IsActive, &u.CreatedBy, &u.UpdatedBy, &u.DeletedBy, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Update(u *entity.User) error {
	query := `
		UPDATE users
		SET first_name=$1, middle_name=$2, surname=$3, age=$4, email=$5, phone_number=$6, role_id=$7, is_active=$8, updated_by=$9, updated_at=NOW()
		WHERE id=$10 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, u.FirstName, u.MiddleName, u.Surname, u.Age, u.Email, u.PhoneNumber, u.RoleID, u.IsActive, u.UpdatedBy, u.ID)
	return err
}

func (r *UserRepository) SoftDelete(id int32, deletedBy int32) error {
	query := `
		UPDATE users
		SET deleted_by=$1, deleted_at=NOW()
		WHERE id=$2 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, deletedBy, id)
	return err
}

func (r *UserRepository) HardDelete(id int32) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
