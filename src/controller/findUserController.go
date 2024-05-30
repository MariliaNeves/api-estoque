package controller

import (
	"github.com/MariliaNeves/api-estoque/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {}

func (uc *userControllerInterface) FindUsers(c *gin.Context) {
	err := rest_err.NewBadRequestError("Erro inexperado, contact adm")
	c.JSON(err.Code, err)
}
