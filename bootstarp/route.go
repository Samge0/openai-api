package bootstarp

import (
	"github.com/gin-gonic/gin"
	"openai-api/routes"
	"sync"
)

var router *gin.Engine
var once sync.Once

func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		routes.RegisterRoutes(router)
	})
}
