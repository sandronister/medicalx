package repository

import (
	"context"

	"github.com/sandronister/medicalx/internal/types/entity"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *entity.Patient) error
	FindByID(ctx context.Context, id int) (*entity.Patient, error)
	FindAll(ctx context.Context) ([]*entity.Patient, error)
	Update(ctx context.Context, patient *entity.Patient) error
	Delete(ctx context.Context, id int) error
}
