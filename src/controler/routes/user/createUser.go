package controler

import (
	"net/http"

	"github.com/MariliaNeves/api-estoque/src/configuration/logger"
	"github.com/MariliaNeves/api-estoque/src/configuration/validation"
	"github.com/MariliaNeves/api-estoque/src/controler/model/request"
	"github.com/MariliaNeves/api-estoque/src/controler/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User create successfully",
		zap.String("journey", "createUser"),
	)
	c.JSON(http.StatusOK, response)
}
