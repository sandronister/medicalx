package usecase

import (
	"context"

	"github.com/sandronister/medicalx/internal/mapper"
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type AnamnesisUseCase struct {
	repo repository.AnamnesisRepository
}

func NewAnamnesisUseCase(repo repository.AnamnesisRepository) *AnamnesisUseCase {
	return &AnamnesisUseCase{repo: repo}
}

func (uc *AnamnesisUseCase) Create(ctx context.Context, req *dto.CreateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
	entity := mapper.CreateAnamnesisRequestToEntity(req)

	if err := uc.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.AnamnesisEntityToResponse(entity), nil
}

func (uc *AnamnesisUseCase) FindByID(ctx context.Context, id int64) (*dto.AnamnesisResponse, error) {
	entity, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.AnamnesisEntityToResponse(entity), nil
}

func (uc *AnamnesisUseCase) FindAll(ctx context.Context) ([]*dto.AnamnesisResponse, error) {
	entities, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.AnamnesisEntitiesToResponses(entities), nil
}

func (uc *AnamnesisUseCase) Update(ctx context.Context, id int64, req *dto.UpdateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
	entity := mapper.UpdateAnamnesisRequestToEntity(id, req)

	if err := uc.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.AnamnesisEntityToResponse(entity), nil
}

func (uc *AnamnesisUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
