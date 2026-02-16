package usecase

import (
	"context"

	"github.com/sandronister/medicalx/internal/mapper"
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type SpecialtyUseCase struct {
	repo repository.SpecialtyRepository
}

func NewSpecialtyUseCase(repo repository.SpecialtyRepository) *SpecialtyUseCase {
	return &SpecialtyUseCase{repo: repo}
}

func (uc *SpecialtyUseCase) Create(ctx context.Context, req *dto.CreateSpecialtyRequest) (*dto.SpecialtyResponse, error) {
	entity := mapper.CreateSpecialtyRequestToEntity(req)

	if err := uc.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.SpecialtyEntityToResponse(entity), nil
}

func (uc *SpecialtyUseCase) FindByID(ctx context.Context, id int64) (*dto.SpecialtyResponse, error) {
	entity, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.SpecialtyEntityToResponse(entity), nil
}

func (uc *SpecialtyUseCase) FindAll(ctx context.Context) ([]*dto.SpecialtyResponse, error) {
	entities, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.SpecialtyEntitiesToResponses(entities), nil
}

func (uc *SpecialtyUseCase) Update(ctx context.Context, id int64, req *dto.UpdateSpecialtyRequest) (*dto.SpecialtyResponse, error) {
	entity := mapper.UpdateSpecialtyRequestToEntity(id, req)

	if err := uc.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.SpecialtyEntityToResponse(entity), nil
}

func (uc *SpecialtyUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
