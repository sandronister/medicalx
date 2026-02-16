package mapper

import (
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/entity"
)

func CreateSpecialtyRequestToEntity(req *dto.CreateSpecialtyRequest) *entity.Specialty {
	return &entity.Specialty{
		Name: req.Name,
	}
}

func UpdateSpecialtyRequestToEntity(id int64, req *dto.UpdateSpecialtyRequest) *entity.Specialty {
	return &entity.Specialty{
		ID:   id,
		Name: req.Name,
	}
}

func SpecialtyEntityToResponse(s *entity.Specialty) *dto.SpecialtyResponse {
	return &dto.SpecialtyResponse{
		ID:   s.ID,
		Name: s.Name,
	}
}

func SpecialtyEntitiesToResponses(specialties []*entity.Specialty) []*dto.SpecialtyResponse {
	responses := make([]*dto.SpecialtyResponse, len(specialties))
	for i, s := range specialties {
		responses[i] = SpecialtyEntityToResponse(s)
	}
	return responses
}
