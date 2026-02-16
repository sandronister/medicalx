package entity

import "time"

type Anamnesis struct {
	ID                    int64
	PatientID             int64
	AppointmentID         int64
	ChiefComplaint        string
	HistoryPresentIllness string
	PastMedicalHistory    string
	FamilyHistory         string
	SocialHistory         string
	ReviewOfSystems       string
	PhysicalExam          string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
