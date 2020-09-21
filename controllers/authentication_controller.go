package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/middleware"
	"go-gin-jwt-authorization-example/models"
	"go-gin-jwt-authorization-example/pkg/app"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/pkg/auth"
	"go-gin-jwt-authorization-example/pkg/gredis"
	"go-gin-jwt-authorization-example/pkg/utils"
	"go-gin-jwt-authorization-example/services"
	"reflect"
	"strconv"

	//_ "github.com/dgrijalva/jwt-go"
)

// Login godoc
// @Summary Login API
// @Tags auth
// @Accept  json
// @Produce  json
// @Param loginDto body dto_request.LoginDto true "loginDto"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/login [post]
func Login(c *gin.Context)  {
	appGin := app.AppGin{C : c}
	var loginDto dto_request.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	var user models.User
	if err := services.FindUserByEmail(&user, loginDto.Email); err != nil {
		appGin.Response(http.StatusNotFound, app_response.USER_NOT_FOUND, nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_USER_EMAIL_OR_PASSWORD, nil)
		return
	}

	tokenDto, err := auth.GenerateToken(strconv.FormatInt(int64(user.Model.ID), 10))
	if err != nil {
		log.Println(err)
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	if err := gredis.Set(utils.GetRedisKey(strconv.FormatInt(int64(user.Model.ID), 10)), tokenDto, tokenDto.AccessTokenExpires); err != nil {
		log.Println(err)
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, tokenDto)
}

// Logout godoc
// @Summary Logout API
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	appGin := app.AppGin{C : c}

	var user models.User
	if err := services.FindUserById(&user, middleware.CurrentUserID); err != nil || reflect.ValueOf(user).IsZero() {
		appGin.Response(http.StatusNotFound, app_response.USER_NOT_FOUND, nil)
		return
	}

	if err := gredis.Delete(utils.GetRedisKey(strconv.Itoa(int(user.Model.ID)))); err !=nil {
		log.Println(err)
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, nil)
}

// RefreshToken godoc
// @Summary Refresh Token API
// @Tags auth
// @Accept  json
// @Produce  json
// @Param refreshTokenRequestDto body dto_request.RefreshTokenRequestDto true "refreshTokenRequestDto"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/refresh-token [post]
func RefreshToken(c *gin.Context) {
	appGin := app.AppGin{C : c}
	var refreshTokenRequestDto dto_request.RefreshTokenRequestDto
	if err := c.ShouldBindJSON(&refreshTokenRequestDto); err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	var parsedToken jwt.Token
	if httpCode, errCode, isError := auth.ParseToken(refreshTokenRequestDto.RefreshToken, &parsedToken); isError {
		appGin.Response(httpCode, errCode, nil)
		return
	}

	var user models.User
	useID , _ := strconv.Atoi(auth.GetUserIDFromToken(parsedToken))

	if err := services.FindUserById(&user, useID); err != nil || reflect.ValueOf(user).IsZero() {
		appGin.Response(http.StatusNotFound, app_response.USER_NOT_FOUND, nil)
		return
	}

	var token auth.TokenDto
	if err := gredis.Get(utils.GetRedisKey(strconv.Itoa(useID)), &token); err != nil {
		log.Println(err)
		appGin.Response(http.StatusNotFound, app_response.TOKEN_NOT_EXISTS, nil)
		return
	}

	if token.RefreshToken != refreshTokenRequestDto.RefreshToken {
		appGin.Response(http.StatusNotFound, app_response.INVALID_TOKEN, nil)
		return
	}

	tokenDto, err := auth.GenerateToken(strconv.FormatInt(int64(user.Model.ID), 10))
	if err != nil {
		log.Println(err)
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	if err := gredis.Set(utils.GetRedisKey(strconv.FormatInt(int64(user.Model.ID), 10)), tokenDto, tokenDto.RefreshTokenExpires); err != nil {
		log.Println(err)
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, tokenDto)
}
