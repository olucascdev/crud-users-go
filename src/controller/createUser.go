package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/olucascdev/crud-users-go/src/configuration/rest_err"
	"github.com/olucascdev/crud-users-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := rest_err.NewBadRequestError("Some fields are incorrect")

		c.JSON(errRest.Code, errRest)
		return

	}
	fmt.Println(userRequest)
}
