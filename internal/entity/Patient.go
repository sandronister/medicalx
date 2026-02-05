package entity

import "time"

type Patient struct {
	ID               int
	FirstName        string
	LastName         string
	DateOfBirth      time.Time
	Gender           string
	Email            string
	Phone            string
	Address          string
	DocumentNumber   string
	EmergencyContact string
	BloodType        string
	Allergies        string
	MedicalHistory   string
	InsuranceNumber  string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
