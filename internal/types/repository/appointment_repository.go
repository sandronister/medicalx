package repository

import (
	"context"

	"github.com/sandronister/medicalx/internal/types/entity"
)

type AppointmentRepository interface {
	Create(ctx context.Context, appointment *entity.Appointment) error
	FindByID(ctx context.Context, id int) (*entity.Appointment, error)
	FindAll(ctx context.Context) ([]*entity.Appointment, error)
	Update(ctx context.Context, appointment *entity.Appointment) error
	Delete(ctx context.Context, id int) error
}
