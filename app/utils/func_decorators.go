package utils

import (
	"fmt"
	"time"
)

// TimeCost 耗时统计函数
func TimeCost(start time.Time) {
	tc := time.Since(start)
	fmt.Printf("time cost = %v\n", tc)
}
