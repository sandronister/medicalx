package dto

type CreatePatientRequest struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DateOfBirth      string `json:"date_of_birth"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	DocumentNumber   string `json:"document_number"`
	EmergencyContact string `json:"emergency_contact"`
	BloodType        string `json:"blood_type"`
	Allergies        string `json:"allergies"`
	MedicalHistory   string `json:"medical_history"`
	InsuranceNumber  string `json:"insurance_number"`
}

type UpdatePatientRequest struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DateOfBirth      string `json:"date_of_birth"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	DocumentNumber   string `json:"document_number"`
	EmergencyContact string `json:"emergency_contact"`
	BloodType        string `json:"blood_type"`
	Allergies        string `json:"allergies"`
	MedicalHistory   string `json:"medical_history"`
	InsuranceNumber  string `json:"insurance_number"`
}

type PatientResponse struct {
	ID               int    `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DateOfBirth      string `json:"date_of_birth"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	DocumentNumber   string `json:"document_number"`
	EmergencyContact string `json:"emergency_contact"`
	BloodType        string `json:"blood_type"`
	Allergies        string `json:"allergies"`
	MedicalHistory   string `json:"medical_history"`
	InsuranceNumber  string `json:"insurance_number"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
