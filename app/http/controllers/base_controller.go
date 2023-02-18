package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (*BaseController) ResponseJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	ctx.Abort()
}
