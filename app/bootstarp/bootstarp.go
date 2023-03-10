package bootstarp

import (
	"openai-api/app/config"
	"openai-api/pkg/logger"
	"strconv"
)

func StartServer() {
	// 注册启动所需各类参数
	SetUpRoute()

	// 启动服务
	port := config.LoadConfig().Port
	portString := strconv.Itoa(port)
	err := router.Run(":" + portString)
	if err != nil {
		logger.Danger("run server error ", err)
		return
	}
}
