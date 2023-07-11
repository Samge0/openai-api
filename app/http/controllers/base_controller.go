package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"openai-api/pkg/logger"
)

type BaseController struct {
}

func (*BaseController) ResponseJson(ctx *gin.Context, code int, msg string, data interface{}) {
	if code != 200 {
		logger.Danger(fmt.Sprintf("[%v]%v", code, msg))
	}
	ctx.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
