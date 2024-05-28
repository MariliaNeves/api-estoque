package service

import (
	"github.com/MariliaNeves/api-estoque/src/configuration/rest_err"
	"github.com/MariliaNeves/api-estoque/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
