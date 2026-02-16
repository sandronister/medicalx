package repository

import (
	"context"
	"database/sql"

	"github.com/sandronister/medicalx/internal/types/entity"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type anamnesisRepositorySQL struct {
	db *sql.DB
}

func NewAnamnesisRepository(db *sql.DB) repository.AnamnesisRepository {
	return &anamnesisRepositorySQL{db: db}
}

func (r *anamnesisRepositorySQL) Create(ctx context.Context, anamnesis *entity.Anamnesis) error {
	query := `INSERT INTO anamnesis 
		(patient_id, appointment_id, chief_complaint, history_present_illness, past_medical_history, 
		 family_history, social_history, review_of_systems, physical_exam, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	result, err := r.db.ExecContext(ctx, query,
		anamnesis.PatientID, anamnesis.AppointmentID, anamnesis.ChiefComplaint,
		anamnesis.HistoryPresentIllness, anamnesis.PastMedicalHistory,
		anamnesis.FamilyHistory, anamnesis.SocialHistory,
		anamnesis.ReviewOfSystems, anamnesis.PhysicalExam,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	anamnesis.ID = id
	return nil
}

func (r *anamnesisRepositorySQL) FindByID(ctx context.Context, id int64) (*entity.Anamnesis, error) {
	query := `SELECT id, patient_id, appointment_id, chief_complaint, history_present_illness, 
		past_medical_history, family_history, social_history, review_of_systems, physical_exam, 
		created_at, updated_at
		FROM anamnesis WHERE id = ?`

	a := &entity.Anamnesis{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.PatientID, &a.AppointmentID, &a.ChiefComplaint,
		&a.HistoryPresentIllness, &a.PastMedicalHistory,
		&a.FamilyHistory, &a.SocialHistory,
		&a.ReviewOfSystems, &a.PhysicalExam,
		&a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *anamnesisRepositorySQL) FindAll(ctx context.Context) ([]*entity.Anamnesis, error) {
	query := `SELECT id, patient_id, appointment_id, chief_complaint, history_present_illness, 
		past_medical_history, family_history, social_history, review_of_systems, physical_exam, 
		created_at, updated_at
		FROM anamnesis`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*entity.Anamnesis
	for rows.Next() {
		a := &entity.Anamnesis{}
		err := rows.Scan(
			&a.ID, &a.PatientID, &a.AppointmentID, &a.ChiefComplaint,
			&a.HistoryPresentIllness, &a.PastMedicalHistory,
			&a.FamilyHistory, &a.SocialHistory,
			&a.ReviewOfSystems, &a.PhysicalExam,
			&a.CreatedAt, &a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}

	return list, rows.Err()
}

func (r *anamnesisRepositorySQL) Update(ctx context.Context, anamnesis *entity.Anamnesis) error {
	query := `UPDATE anamnesis SET 
		patient_id = ?, appointment_id = ?, chief_complaint = ?, history_present_illness = ?, 
		past_medical_history = ?, family_history = ?, social_history = ?, review_of_systems = ?, 
		physical_exam = ?, updated_at = NOW()
		WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		anamnesis.PatientID, anamnesis.AppointmentID, anamnesis.ChiefComplaint,
		anamnesis.HistoryPresentIllness, anamnesis.PastMedicalHistory,
		anamnesis.FamilyHistory, anamnesis.SocialHistory,
		anamnesis.ReviewOfSystems, anamnesis.PhysicalExam,
		anamnesis.ID,
	)
	return err
}

func (r *anamnesisRepositorySQL) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM anamnesis WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
