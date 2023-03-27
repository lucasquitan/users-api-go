package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasquitan/users-api-go/src/configuration/logger"
	"github.com/lucasquitan/users-api-go/src/configuration/validation"
	"github.com/lucasquitan/users-api-go/src/controller/model/request"
	"github.com/lucasquitan/users-api-go/src/controller/model/request/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init the CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userResquest.Email,
		Name:  userResquest.Name,
		Age:   int8(userResquest.Age),
	}

	logger.Info("User created successgully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, response)
}
