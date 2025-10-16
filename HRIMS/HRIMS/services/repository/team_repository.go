package repository

import (
	"database/sql"
	"training-backend/services/entity"
	"training-backend/package/log"
)

type TeamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Create(team *entity.Team) error {
	query := `
		INSERT INTO teams (team_name, description, created_by, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id`
	return r.db.QueryRow(query, team.TeamName, team.Description, team.CreatedBy).Scan(&team.ID)
}

func (r *TeamRepository) List() ([]*entity.Team, error) {
	query := `
		SELECT id, team_name, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM teams
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("error fetching teams: %v", err)
		return nil, err
	}
	defer rows.Close()

	var teams []*entity.Team
	for rows.Next() {
		t := &entity.Team{}
		err := rows.Scan(&t.ID, &t.TeamName, &t.Description, &t.CreatedBy, &t.UpdatedBy, &t.DeletedBy, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
		if err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	return teams, nil
}

func (r *TeamRepository) Get(id int32) (*entity.Team, error) {
	query := `
		SELECT id, team_name, description, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at
		FROM teams
		WHERE id = $1 AND deleted_at IS NULL`
	t := &entity.Team{}
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.TeamName, &t.Description, &t.CreatedBy, &t.UpdatedBy, &t.DeletedBy, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TeamRepository) Update(t *entity.Team) error {
	query := `
		UPDATE teams
		SET team_name = $1, description = $2, updated_by = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, t.TeamName, t.Description, t.UpdatedBy, t.ID)
	return err
}

func (r *TeamRepository) SoftDelete(id int32, deletedBy int32) error {
	query := `
		UPDATE teams
		SET deleted_by = $1, deleted_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, deletedBy, id)
	return err
}

func (r *TeamRepository) HardDelete(id int32) error {
	_, err := r.db.Exec("DELETE FROM teams WHERE id=$1", id)
	return err
}
