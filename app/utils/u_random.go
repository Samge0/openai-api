package utils

import (
	"math/rand"
	"openai-api/app/config"
	"openai-api/pkg/logger"
	"strings"
	"time"
)

// GetRandomApiKey 获取随机的openai的token值
func GetRandomApiKey() string {
	apiKey := config.LoadConfig().ApiKey
	if strings.Contains(apiKey, "|") {
		apiKeys := strings.Split(apiKey, "|")
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(apiKeys))
		apiKey = apiKeys[randomIndex]
	}
	logger.Info("GetRandomApiKey apiKey is ", apiKey)
	return apiKey
}
