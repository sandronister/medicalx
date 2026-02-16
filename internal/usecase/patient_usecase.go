package usecase

import (
	"context"

	"github.com/sandronister/medicalx/internal/mapper"
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type PatientUseCase struct {
	repo repository.PatientRepository
}

func NewPatientUseCase(repo repository.PatientRepository) *PatientUseCase {
	return &PatientUseCase{repo: repo}
}

func (uc *PatientUseCase) Create(ctx context.Context, req *dto.CreatePatientRequest) (*dto.PatientResponse, error) {
	entity, err := mapper.CreatePatientRequestToEntity(req)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.PatientEntityToResponse(entity), nil
}

func (uc *PatientUseCase) FindByID(ctx context.Context, id int) (*dto.PatientResponse, error) {
	entity, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.PatientEntityToResponse(entity), nil
}

func (uc *PatientUseCase) FindAll(ctx context.Context) ([]*dto.PatientResponse, error) {
	entities, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.PatientEntitiesToResponses(entities), nil
}

func (uc *PatientUseCase) Update(ctx context.Context, id int, req *dto.UpdatePatientRequest) (*dto.PatientResponse, error) {
	entity, err := mapper.UpdatePatientRequestToEntity(id, req)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.PatientEntityToResponse(entity), nil
}

func (uc *PatientUseCase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
