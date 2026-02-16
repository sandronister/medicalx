package dto

type CreateAppointmentRequest struct {
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
	Status    string `json:"status"`
	Notes     string `json:"notes"`
	PatientID int    `json:"patient_id"`
	MedicoID  int    `json:"medico_id"`
}

type UpdateAppointmentRequest struct {
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
	Status   string `json:"status"`
	Notes    string `json:"notes"`
}

type AppointmentResponse struct {
	ID        int    `json:"id"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
	Status    string `json:"status"`
	Notes     string `json:"notes"`
	PatientID int    `json:"patient_id"`
	MedicoID  int    `json:"medico_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
