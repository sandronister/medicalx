package usecase

import (
	"context"

	"github.com/sandronister/medicalx/internal/mapper"
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	entity := mapper.CreateUserRequestToEntity(req)

	if err := uc.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.UserEntityToResponse(entity), nil
}

func (uc *UserUseCase) FindByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	entity, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.UserEntityToResponse(entity), nil
}

func (uc *UserUseCase) FindAll(ctx context.Context) ([]*dto.UserResponse, error) {
	entities, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.UserEntitiesToResponses(entities), nil
}

func (uc *UserUseCase) Update(ctx context.Context, id int, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	entity := mapper.UpdateUserRequestToEntity(id, req)

	if err := uc.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return mapper.UserEntityToResponse(entity), nil
}

func (uc *UserUseCase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
