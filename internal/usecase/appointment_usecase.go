package usecase

import (
	"context"

	"github.com/sandronister/medicalx/internal/mapper"
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type AppointmentUseCase struct {
	repo repository.AppointmentRepository
}

func NewAppointmentUseCase(repo repository.AppointmentRepository) *AppointmentUseCase {
	return &AppointmentUseCase{repo: repo}
}

func (uc *AppointmentUseCase) Create(ctx context.Context, req *dto.CreateAppointmentRequest) (*dto.AppointmentResponse, error) {
	entity := mapper.CreateAppointmentRequestToEntity(req)

	if err := uc.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.AppointmentEntityToResponse(entity), nil
}

func (uc *AppointmentUseCase) FindByID(ctx context.Context, id int) (*dto.AppointmentResponse, error) {
	entity, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.AppointmentEntityToResponse(entity), nil
}

func (uc *AppointmentUseCase) FindAll(ctx context.Context) ([]*dto.AppointmentResponse, error) {
	entities, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.AppointmentEntitiesToResponses(entities), nil
}

func (uc *AppointmentUseCase) Update(ctx context.Context, id int, req *dto.UpdateAppointmentRequest) (*dto.AppointmentResponse, error) {
	entity := mapper.UpdateAppointmentRequestToEntity(id, req)

	if err := uc.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.AppointmentEntityToResponse(entity), nil
}

func (uc *AppointmentUseCase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
