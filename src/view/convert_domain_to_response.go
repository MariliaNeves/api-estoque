package view

import (
	"github.com/MariliaNeves/api-estoque/src/controller/model/response"
	"github.com/MariliaNeves/api-estoque/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
