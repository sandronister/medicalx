package server

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	di "github.com/sandronister/medicalx/internal/DI"
)

type Server struct {
	app *fiber.App
	db  *sql.DB
}

func NewServer(db *sql.DB) *Server {
	app := fiber.New(fiber.Config{
		AppName: "MedicalX",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	s := &Server{
		app: app,
		db:  db,
	}

	s.registerRoutes()

	return s
}

func (s *Server) registerRoutes() {
	api := s.app.Group("/api")

	// Patient
	patientHandler := di.NewPatientHandler(s.db)

	patients := api.Group("/patients")
	patients.Post("/", patientHandler.Create)
	patients.Get("/", patientHandler.FindAll)
	patients.Get("/:id", patientHandler.FindByID)
	patients.Put("/:id", patientHandler.Update)
	patients.Delete("/:id", patientHandler.Delete)

	// User
	userHandler := di.NewUserHandler(s.db)

	users := api.Group("/users")
	users.Post("/", userHandler.Create)
	users.Get("/", userHandler.FindAll)
	users.Get("/:id", userHandler.FindByID)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)

	// Appointment
	appointmentHandler := di.NewAppointmentHandler(s.db)

	appointments := api.Group("/appointments")
	appointments.Post("/", appointmentHandler.Create)
	appointments.Get("/", appointmentHandler.FindAll)
	appointments.Get("/:id", appointmentHandler.FindByID)
	appointments.Put("/:id", appointmentHandler.Update)
	appointments.Delete("/:id", appointmentHandler.Delete)

	// Anamnesis
	anamnesisHandler := di.NewAnamnesisHandler(s.db)

	anamnesis := api.Group("/anamnesis")
	anamnesis.Post("/", anamnesisHandler.Create)
	anamnesis.Get("/", anamnesisHandler.FindAll)
	anamnesis.Get("/:id", anamnesisHandler.FindByID)
	anamnesis.Put("/:id", anamnesisHandler.Update)
	anamnesis.Delete("/:id", anamnesisHandler.Delete)

	// Specialty
	specialtyHandler := di.NewSpecialtyHandler(s.db)

	specialties := api.Group("/specialties")
	specialties.Post("/", specialtyHandler.Create)
	specialties.Get("/", specialtyHandler.FindAll)
	specialties.Get("/:id", specialtyHandler.FindByID)
	specialties.Put("/:id", specialtyHandler.Update)
	specialties.Delete("/:id", specialtyHandler.Delete)
}

func (s *Server) Start(addr string) error {
	log.Printf("Server starting on %s", addr)
	return s.app.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
