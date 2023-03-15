package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (*BaseController) ResponseJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
