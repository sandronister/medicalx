package mapper

import (
	"github.com/sandronister/medicalx/internal/types/dto"
	"github.com/sandronister/medicalx/internal/types/entity"
)

func CreateUserRequestToEntity(req *dto.CreateUserRequest) *entity.User {
	return &entity.User{
		Name:      req.Name,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		Celular:   req.Celular,
		Whatsapp:  req.Whatsapp,
		Nif:       req.Nif,
		Age:       req.Age,
		Gender:    req.Gender,
		Crm:       req.Crm,
		CompanyID: req.CompanyID,
	}
}

func UpdateUserRequestToEntity(id int, req *dto.UpdateUserRequest) *entity.User {
	return &entity.User{
		ID:       id,
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Role:     req.Role,
		Celular:  req.Celular,
		Whatsapp: req.Whatsapp,
		Nif:      req.Nif,
		Age:      req.Age,
		Gender:   req.Gender,
		Crm:      req.Crm,
	}
}

func UserEntityToResponse(u *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Role:      u.Role,
		Celular:   u.Celular,
		Whatsapp:  u.Whatsapp,
		Nif:       u.Nif,
		Age:       u.Age,
		Gender:    u.Gender,
		Crm:       u.Crm,
		CompanyID: u.CompanyID,
	}
}

func UserEntitiesToResponses(users []*entity.User) []*dto.UserResponse {
	responses := make([]*dto.UserResponse, len(users))
	for i, u := range users {
		responses[i] = UserEntityToResponse(u)
	}
	return responses
}
