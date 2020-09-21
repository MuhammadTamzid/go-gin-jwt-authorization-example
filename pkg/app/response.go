package app

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
)

type AppGin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//Response setting gin.JSON
func (g *AppGin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errCode,
		Message: app_response.GetMsg(errCode),
		Data:    data,
	})
	return
}
