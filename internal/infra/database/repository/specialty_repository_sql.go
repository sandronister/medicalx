package repository

import (
	"context"
	"database/sql"

	"github.com/sandronister/medicalx/internal/types/entity"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type specialtyRepositorySQL struct {
	db *sql.DB
}

func NewSpecialtyRepository(db *sql.DB) repository.SpecialtyRepository {
	return &specialtyRepositorySQL{db: db}
}

func (r *specialtyRepositorySQL) Create(ctx context.Context, specialty *entity.Specialty) error {
	query := `INSERT INTO specialties (name) VALUES (?)`

	result, err := r.db.ExecContext(ctx, query, specialty.Name)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	specialty.ID = id
	return nil
}

func (r *specialtyRepositorySQL) FindByID(ctx context.Context, id int64) (*entity.Specialty, error) {
	query := `SELECT id, name FROM specialties WHERE id = ?`

	s := &entity.Specialty{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&s.ID, &s.Name)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (r *specialtyRepositorySQL) FindAll(ctx context.Context) ([]*entity.Specialty, error) {
	query := `SELECT id, name FROM specialties`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specialties []*entity.Specialty
	for rows.Next() {
		s := &entity.Specialty{}
		err := rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return nil, err
		}
		specialties = append(specialties, s)
	}

	return specialties, rows.Err()
}

func (r *specialtyRepositorySQL) Update(ctx context.Context, specialty *entity.Specialty) error {
	query := `UPDATE specialties SET name = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, specialty.Name, specialty.ID)
	return err
}

func (r *specialtyRepositorySQL) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM specialties WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
