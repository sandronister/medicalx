package repository

import (
	"context"
	"database/sql"

	"github.com/sandronister/medicalx/internal/types/entity"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type userRepositorySQL struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositorySQL{db: db}
}

func (r *userRepositorySQL) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users
		(name, surname, email, password, role, celular, whatsapp, nif, age, gender, crm, company_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query,
		user.Name, user.Surname, user.Email, user.Password, user.Role,
		user.Celular, user.Whatsapp, user.Nif, user.Age, user.Gender,
		user.Crm, user.CompanyID,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *userRepositorySQL) FindByID(ctx context.Context, id int) (*entity.User, error) {
	query := `SELECT id, name, surname, email, password, role, celular, whatsapp, nif, age, gender, crm, company_id
		FROM users WHERE id = ?`

	u := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID, &u.Name, &u.Surname, &u.Email, &u.Password, &u.Role,
		&u.Celular, &u.Whatsapp, &u.Nif, &u.Age, &u.Gender,
		&u.Crm, &u.CompanyID,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepositorySQL) FindAll(ctx context.Context) ([]*entity.User, error) {
	query := `SELECT id, name, surname, email, password, role, celular, whatsapp, nif, age, gender, crm, company_id
		FROM users`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		u := &entity.User{}
		err := rows.Scan(
			&u.ID, &u.Name, &u.Surname, &u.Email, &u.Password, &u.Role,
			&u.Celular, &u.Whatsapp, &u.Nif, &u.Age, &u.Gender,
			&u.Crm, &u.CompanyID,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *userRepositorySQL) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET
		name = ?, surname = ?, email = ?, role = ?, celular = ?, whatsapp = ?,
		nif = ?, age = ?, gender = ?, crm = ?
		WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		user.Name, user.Surname, user.Email, user.Role,
		user.Celular, user.Whatsapp, user.Nif, user.Age, user.Gender,
		user.Crm, user.ID,
	)
	return err
}

func (r *userRepositorySQL) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
