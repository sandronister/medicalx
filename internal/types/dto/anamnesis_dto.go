package dto

type CreateAnamnesisRequest struct {
	PatientID             int64  `json:"patient_id"`
	AppointmentID         int64  `json:"appointment_id"`
	ChiefComplaint        string `json:"chief_complaint"`
	HistoryPresentIllness string `json:"history_present_illness"`
	PastMedicalHistory    string `json:"past_medical_history"`
	FamilyHistory         string `json:"family_history"`
	SocialHistory         string `json:"social_history"`
	ReviewOfSystems       string `json:"review_of_systems"`
	PhysicalExam          string `json:"physical_exam"`
}

type UpdateAnamnesisRequest struct {
	ChiefComplaint        string `json:"chief_complaint"`
	HistoryPresentIllness string `json:"history_present_illness"`
	PastMedicalHistory    string `json:"past_medical_history"`
	FamilyHistory         string `json:"family_history"`
	SocialHistory         string `json:"social_history"`
	ReviewOfSystems       string `json:"review_of_systems"`
	PhysicalExam          string `json:"physical_exam"`
}

type AnamnesisResponse struct {
	ID                    int64  `json:"id"`
	PatientID             int64  `json:"patient_id"`
	AppointmentID         int64  `json:"appointment_id"`
	ChiefComplaint        string `json:"chief_complaint"`
	HistoryPresentIllness string `json:"history_present_illness"`
	PastMedicalHistory    string `json:"past_medical_history"`
	FamilyHistory         string `json:"family_history"`
	SocialHistory         string `json:"social_history"`
	ReviewOfSystems       string `json:"review_of_systems"`
	PhysicalExam          string `json:"physical_exam"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}
