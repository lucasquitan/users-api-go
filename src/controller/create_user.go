package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucasquitan/users-api-go/src/configuration/validation"
	"github.com/lucasquitan/users-api-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init the CreateUser controller")
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		log.Printf("Error trying to marshal objectm error=%s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userResquest)
}
