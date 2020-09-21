package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-gin-jwt-authorization-example/models"
	"go-gin-jwt-authorization-example/pkg/app"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/services"
)

type AppAuthorizer struct {
	enforcer *casbin.Enforcer
}

func Authorizer(e *casbin.Enforcer) gin.HandlerFunc {
	authz := &AppAuthorizer{enforcer: e}
	return func(c *gin.Context) {
		appGin := app.AppGin{C : c}
		if !authz.isAllowedPermission(c.Request) {
			appGin.Response(http.StatusForbidden, app_response.NOT_AUTHORIZED, nil)
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

func (a *AppAuthorizer) isAllowedPermission(r *http.Request) bool {
	var user models.User
	method := r.Method
	path := r.URL.Path

	if err := services.FindUserById(&user, CurrentUserID); err != nil {
		return false
	}

	isAllowed, err := a.enforcer.Enforce(user.Role.Name, path, method)
	if err != nil {
		panic(err)
	}

	return isAllowed
}
