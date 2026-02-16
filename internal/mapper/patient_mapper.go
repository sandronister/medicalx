package mapper

import (
	"time"

	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/entity"
)

const dateLayout = "2006-01-02"
const dateTimeLayout = "2006-01-02T15:04:05Z"

func CreatePatientRequestToEntity(req *dto.CreatePatientRequest) (*entity.Patient, error) {
	dob, err := time.Parse(dateLayout, req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &entity.Patient{
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		DateOfBirth:      dob,
		Gender:           req.Gender,
		Email:            req.Email,
		Phone:            req.Phone,
		Address:          req.Address,
		DocumentNumber:   req.DocumentNumber,
		EmergencyContact: req.EmergencyContact,
		BloodType:        req.BloodType,
		Allergies:        req.Allergies,
		MedicalHistory:   req.MedicalHistory,
		InsuranceNumber:  req.InsuranceNumber,
	}, nil
}

func UpdatePatientRequestToEntity(id int, req *dto.UpdatePatientRequest) (*entity.Patient, error) {
	dob, err := time.Parse(dateLayout, req.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &entity.Patient{
		ID:               id,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		DateOfBirth:      dob,
		Gender:           req.Gender,
		Email:            req.Email,
		Phone:            req.Phone,
		Address:          req.Address,
		DocumentNumber:   req.DocumentNumber,
		EmergencyContact: req.EmergencyContact,
		BloodType:        req.BloodType,
		Allergies:        req.Allergies,
		MedicalHistory:   req.MedicalHistory,
		InsuranceNumber:  req.InsuranceNumber,
	}, nil
}

func PatientEntityToResponse(p *entity.Patient) *dto.PatientResponse {
	return &dto.PatientResponse{
		ID:               p.ID,
		FirstName:        p.FirstName,
		LastName:         p.LastName,
		DateOfBirth:      p.DateOfBirth.Format(dateLayout),
		Gender:           p.Gender,
		Email:            p.Email,
		Phone:            p.Phone,
		Address:          p.Address,
		DocumentNumber:   p.DocumentNumber,
		EmergencyContact: p.EmergencyContact,
		BloodType:        p.BloodType,
		Allergies:        p.Allergies,
		MedicalHistory:   p.MedicalHistory,
		InsuranceNumber:  p.InsuranceNumber,
		CreatedAt:        p.CreatedAt.Format(dateTimeLayout),
		UpdatedAt:        p.UpdatedAt.Format(dateTimeLayout),
	}
}

func PatientEntitiesToResponses(patients []*entity.Patient) []*dto.PatientResponse {
	responses := make([]*dto.PatientResponse, len(patients))
	for i, p := range patients {
		responses[i] = PatientEntityToResponse(p)
	}
	return responses
}
