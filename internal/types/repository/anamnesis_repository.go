package repository

import (
	"context"

	"github.com/sandronister/medicalx/internal/types/entity"
)

type AnamnesisRepository interface {
	Create(ctx context.Context, anamnesis *entity.Anamnesis) error
	FindByID(ctx context.Context, id int64) (*entity.Anamnesis, error)
	FindAll(ctx context.Context) ([]*entity.Anamnesis, error)
	Update(ctx context.Context, anamnesis *entity.Anamnesis) error
	Delete(ctx context.Context, id int64) error
}
