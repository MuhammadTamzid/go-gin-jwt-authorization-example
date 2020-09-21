package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/pkg/app"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/services"
)

// Register godoc
// @Summary User Registration API
// @Tags users
// @Accept  json
// @Produce  json
// @Param userRegistrationDto body dto_request.UserRegistrationDto true "userRegistrationDto"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router /users/register [post]
func Register(c *gin.Context)  {
	appGin := app.AppGin{C :c}
	var userRegistrationDto dto_request.UserRegistrationDto
	if err := c.ShouldBindJSON(&userRegistrationDto); err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	if err := services.RegisterUser(userRegistrationDto); err != nil {
		appGin.Response(http.StatusBadRequest, app_response.FAIL_TO_REGISTER_USER, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, nil)
}
