package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-gin-jwt-authorization-example/models"
	"go-gin-jwt-authorization-example/pkg/app"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/pkg/auth"
	"go-gin-jwt-authorization-example/pkg/constant"
	"go-gin-jwt-authorization-example/pkg/gredis"
	"go-gin-jwt-authorization-example/pkg/utils"
	"go-gin-jwt-authorization-example/services"
	"strconv"
)

// TODO: Refactor
var CurrentUserID int

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appGin := app.AppGin{C : c}

		var tokenString string
		if httpCode, errCode, isError := auth.GetTokenFromHeader(c.GetHeader(constant.HEADER_STRING_AUTHORIZATION), &tokenString); isError {
			appGin.Response(httpCode, errCode, nil)
			c.Abort()
			return
		}

		var parsedToken jwt.Token
		if httpCode, errCode, isError := auth.ParseToken(tokenString, &parsedToken); isError {
			appGin.Response(httpCode, errCode, nil)
			c.Abort()
			return
		}

		var user models.User
		CurrentUserID , _ = strconv.Atoi(auth.GetUserIDFromToken(parsedToken))

		if err := services.FindUserById(&user, CurrentUserID); err != nil {
			appGin.Response(http.StatusNotFound, app_response.USER_NOT_FOUND, nil)
			c.Abort()
			return
		}

		var token auth.TokenDto
		if err := gredis.Get(utils.GetRedisKey(strconv.Itoa(CurrentUserID)), &token); err != nil {
			appGin.Response(http.StatusNotFound, app_response.TOKEN_NOT_EXISTS, nil)
			c.Abort()
			return
		}

		if token.AccessToken != tokenString {
			appGin.Response(http.StatusNotFound, app_response.INVALID_TOKEN, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
