package routes

import (
	"github.com/gin-gonic/gin"
	. "openai-api/app/http/controllers"
	"openai-api/app/middlewares"
)

var chatController = NewChatController()

// RegisterRoutes 注册路由
func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	apiRouter := router.Group("api").Use(middlewares.TokenJWTAuth())
	{
		apiRouter.POST("/chat", chatController.HandlerChat)
	}
}
