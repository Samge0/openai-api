package utils

import (
	"openai-api/pkg/logger"
	"time"
)

// TimeCost 耗时统计函数
func TimeCost(start time.Time) {
	tc := time.Since(start)
	logger.Info("time cost = ", tc)
}
