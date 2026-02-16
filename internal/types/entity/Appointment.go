package entity

type Appointment struct {
	ID        int
	StartsAt  string
	EndsAt    string
	Status    string
	Notes     string
	PatientID int
	MedicoID  int
	CreatedAt string
	UpdatedAt string
}
