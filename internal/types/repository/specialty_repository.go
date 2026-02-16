package repository

import (
	"context"

	"github.com/sandronister/medicalx/internal/types/entity"
)

type SpecialtyRepository interface {
	Create(ctx context.Context, specialty *entity.Specialty) error
	FindByID(ctx context.Context, id int64) (*entity.Specialty, error)
	FindAll(ctx context.Context) ([]*entity.Specialty, error)
	Update(ctx context.Context, specialty *entity.Specialty) error
	Delete(ctx context.Context, id int64) error
}
