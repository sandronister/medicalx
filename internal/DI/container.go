package di

import (
	"database/sql"

	"github.com/sandronister/medicalx/internal/infra/database/repository"
	"github.com/sandronister/medicalx/internal/infra/web/handler"
	"github.com/sandronister/medicalx/internal/usecase"
)

func NewPatientHandler(db *sql.DB) *handler.PatientHandler {
	repo := repository.NewPatientRepository(db)
	uc := usecase.NewPatientUseCase(repo)
	return handler.NewPatientHandler(uc)
}

func NewUserHandler(db *sql.DB) *handler.UserHandler {
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	return handler.NewUserHandler(uc)
}

func NewAppointmentHandler(db *sql.DB) *handler.AppointmentHandler {
	repo := repository.NewAppointmentRepository(db)
	uc := usecase.NewAppointmentUseCase(repo)
	return handler.NewAppointmentHandler(uc)
}

func NewAnamnesisHandler(db *sql.DB) *handler.AnamnesisHandler {
	repo := repository.NewAnamnesisRepository(db)
	uc := usecase.NewAnamnesisUseCase(repo)
	return handler.NewAnamnesisHandler(uc)
}

func NewSpecialtyHandler(db *sql.DB) *handler.SpecialtyHandler {
	repo := repository.NewSpecialtyRepository(db)
	uc := usecase.NewSpecialtyUseCase(repo)
	return handler.NewSpecialtyHandler(uc)
}