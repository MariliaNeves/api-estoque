package controler

import (
	"github.com/MariliaNeves/api-estoque/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {}

func FindUsers(c *gin.Context) {
	err := rest_err.NewBadRequestError("Erro inexperado, contact adm")
	c.JSON(err.Code, err)
}
