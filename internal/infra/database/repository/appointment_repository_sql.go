package repository

import (
	"context"
	"database/sql"

	"github.com/sandronister/medicalx/internal/types/entity"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type appointmentRepositorySQL struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) repository.AppointmentRepository {
	return &appointmentRepositorySQL{db: db}
}

func (r *appointmentRepositorySQL) Create(ctx context.Context, appointment *entity.Appointment) error {
	query := `INSERT INTO appointments
		(starts_at, ends_at, status, notes, patient_id, medico_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`

	result, err := r.db.ExecContext(ctx, query,
		appointment.StartsAt, appointment.EndsAt, appointment.Status,
		appointment.Notes, appointment.PatientID, appointment.MedicoID,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	appointment.ID = int(id)
	return nil
}

func (r *appointmentRepositorySQL) FindByID(ctx context.Context, id int) (*entity.Appointment, error) {
	query := `SELECT id, starts_at, ends_at, status, notes, patient_id, medico_id, created_at, updated_at
		FROM appointments WHERE id = ?`

	a := &entity.Appointment{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.StartsAt, &a.EndsAt, &a.Status,
		&a.Notes, &a.PatientID, &a.MedicoID, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *appointmentRepositorySQL) FindAll(ctx context.Context) ([]*entity.Appointment, error) {
	query := `SELECT id, starts_at, ends_at, status, notes, patient_id, medico_id, created_at, updated_at
		FROM appointments`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []*entity.Appointment
	for rows.Next() {
		a := &entity.Appointment{}
		err := rows.Scan(
			&a.ID, &a.StartsAt, &a.EndsAt, &a.Status,
			&a.Notes, &a.PatientID, &a.MedicoID, &a.CreatedAt, &a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, a)
	}

	return appointments, rows.Err()
}

func (r *appointmentRepositorySQL) Update(ctx context.Context, appointment *entity.Appointment) error {
	query := `UPDATE appointments SET
		starts_at = ?, ends_at = ?, status = ?, notes = ?, patient_id = ?, medico_id = ?, updated_at = NOW()
		WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		appointment.StartsAt, appointment.EndsAt, appointment.Status,
		appointment.Notes, appointment.PatientID, appointment.MedicoID,
		appointment.ID,
	)
	return err
}

func (r *appointmentRepositorySQL) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM appointments WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
