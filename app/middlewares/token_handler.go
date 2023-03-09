package middlewares

import (
	"github.com/gin-gonic/gin"
	"openai-api/app/config"
	"strings"
)

// TokenJWTAuth 拦截器
func TokenJWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 如果没配置，则直接通过
		if len(config.LoadConfig().AccessToken) == 0 {
			ctx.Next()
			return
		}

		// 这里做简单校验
		token := ctx.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		if token != config.LoadConfig().AccessToken {
			ResponseJson(ctx, 403, "非法访问", nil)
			return
		}
		ctx.Next()
	}

}

func ResponseJson(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
	if code != 200 {
		ctx.Abort()
	}
}
