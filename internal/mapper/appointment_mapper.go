package mapper

import (
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/entity"
)

func CreateAppointmentRequestToEntity(req *dto.CreateAppointmentRequest) *entity.Appointment {
	return &entity.Appointment{
		StartsAt:  req.StartsAt,
		EndsAt:    req.EndsAt,
		Status:    req.Status,
		Notes:     req.Notes,
		PatientID: req.PatientID,
		MedicoID:  req.MedicoID,
	}
}

func UpdateAppointmentRequestToEntity(id int, req *dto.UpdateAppointmentRequest) *entity.Appointment {
	return &entity.Appointment{
		ID:       id,
		StartsAt: req.StartsAt,
		EndsAt:   req.EndsAt,
		Status:   req.Status,
		Notes:    req.Notes,
	}
}

func AppointmentEntityToResponse(a *entity.Appointment) *dto.AppointmentResponse {
	return &dto.AppointmentResponse{
		ID:        a.ID,
		StartsAt:  a.StartsAt,
		EndsAt:    a.EndsAt,
		Status:    a.Status,
		Notes:     a.Notes,
		PatientID: a.PatientID,
		MedicoID:  a.MedicoID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func AppointmentEntitiesToResponses(appointments []*entity.Appointment) []*dto.AppointmentResponse {
	responses := make([]*dto.AppointmentResponse, len(appointments))
	for i, a := range appointments {
		responses[i] = AppointmentEntityToResponse(a)
	}
	return responses
}
