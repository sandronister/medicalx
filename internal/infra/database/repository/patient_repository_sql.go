package repository

import (
	"context"
	"database/sql"

	"github.com/sandronister/medicalx/internal/types/entity"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type patientRepositorySQL struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) repository.PatientRepository {
	return &patientRepositorySQL{db: db}
}

func (r *patientRepositorySQL) Create(ctx context.Context, patient *entity.Patient) error {
	query := `INSERT INTO patients
		(first_name, last_name, date_of_birth, gender, email, phone, address,
		 document_number, emergency_contact, blood_type, allergies, medical_history, insurance_number, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	result, err := r.db.ExecContext(ctx, query,
		patient.FirstName, patient.LastName, patient.DateOfBirth, patient.Gender,
		patient.Email, patient.Phone, patient.Address, patient.DocumentNumber,
		patient.EmergencyContact, patient.BloodType, patient.Allergies,
		patient.MedicalHistory, patient.InsuranceNumber,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	patient.ID = int(id)
	return nil
}

func (r *patientRepositorySQL) FindByID(ctx context.Context, id int) (*entity.Patient, error) {
	query := `SELECT id, first_name, last_name, date_of_birth, gender, email, phone, address,
		document_number, emergency_contact, blood_type, allergies, medical_history, insurance_number,
		created_at, updated_at
		FROM patients WHERE id = ?`

	p := &entity.Patient{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth, &p.Gender,
		&p.Email, &p.Phone, &p.Address, &p.DocumentNumber,
		&p.EmergencyContact, &p.BloodType, &p.Allergies,
		&p.MedicalHistory, &p.InsuranceNumber, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *patientRepositorySQL) FindAll(ctx context.Context) ([]*entity.Patient, error) {
	query := `SELECT id, first_name, last_name, date_of_birth, gender, email, phone, address,
		document_number, emergency_contact, blood_type, allergies, medical_history, insurance_number,
		created_at, updated_at
		FROM patients`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []*entity.Patient
	for rows.Next() {
		p := &entity.Patient{}
		err := rows.Scan(
			&p.ID, &p.FirstName, &p.LastName, &p.DateOfBirth, &p.Gender,
			&p.Email, &p.Phone, &p.Address, &p.DocumentNumber,
			&p.EmergencyContact, &p.BloodType, &p.Allergies,
			&p.MedicalHistory, &p.InsuranceNumber, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, rows.Err()
}

func (r *patientRepositorySQL) Update(ctx context.Context, patient *entity.Patient) error {
	query := `UPDATE patients SET
		first_name = ?, last_name = ?, date_of_birth = ?, gender = ?, email = ?, phone = ?,
		address = ?, document_number = ?, emergency_contact = ?, blood_type = ?, allergies = ?,
		medical_history = ?, insurance_number = ?, updated_at = NOW()
		WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		patient.FirstName, patient.LastName, patient.DateOfBirth, patient.Gender,
		patient.Email, patient.Phone, patient.Address, patient.DocumentNumber,
		patient.EmergencyContact, patient.BloodType, patient.Allergies,
		patient.MedicalHistory, patient.InsuranceNumber, patient.ID,
	)
	return err
}

func (r *patientRepositorySQL) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM patients WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
