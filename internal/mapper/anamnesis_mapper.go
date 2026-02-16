package mapper

import (
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/entity"
)

func CreateAnamnesisRequestToEntity(req *dto.CreateAnamnesisRequest) *entity.Anamnesis {
	return &entity.Anamnesis{
		PatientID:             req.PatientID,
		AppointmentID:         req.AppointmentID,
		ChiefComplaint:        req.ChiefComplaint,
		HistoryPresentIllness: req.HistoryPresentIllness,
		PastMedicalHistory:    req.PastMedicalHistory,
		FamilyHistory:         req.FamilyHistory,
		SocialHistory:         req.SocialHistory,
		ReviewOfSystems:       req.ReviewOfSystems,
		PhysicalExam:          req.PhysicalExam,
	}
}

func UpdateAnamnesisRequestToEntity(id int64, req *dto.UpdateAnamnesisRequest) *entity.Anamnesis {
	return &entity.Anamnesis{
		ID:                    id,
		ChiefComplaint:        req.ChiefComplaint,
		HistoryPresentIllness: req.HistoryPresentIllness,
		PastMedicalHistory:    req.PastMedicalHistory,
		FamilyHistory:         req.FamilyHistory,
		SocialHistory:         req.SocialHistory,
		ReviewOfSystems:       req.ReviewOfSystems,
		PhysicalExam:          req.PhysicalExam,
	}
}

func AnamnesisEntityToResponse(a *entity.Anamnesis) *dto.AnamnesisResponse {
	return &dto.AnamnesisResponse{
		ID:                    a.ID,
		PatientID:             a.PatientID,
		AppointmentID:         a.AppointmentID,
		ChiefComplaint:        a.ChiefComplaint,
		HistoryPresentIllness: a.HistoryPresentIllness,
		PastMedicalHistory:    a.PastMedicalHistory,
		FamilyHistory:         a.FamilyHistory,
		SocialHistory:         a.SocialHistory,
		ReviewOfSystems:       a.ReviewOfSystems,
		PhysicalExam:          a.PhysicalExam,
		CreatedAt:             a.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:             a.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func AnamnesisEntitiesToResponses(list []*entity.Anamnesis) []*dto.AnamnesisResponse {
	responses := make([]*dto.AnamnesisResponse, len(list))
	for i, a := range list {
		responses[i] = AnamnesisEntityToResponse(a)
	}
	return responses
}
