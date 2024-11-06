package secondModul

import (
	"fmt"
	"go.uber.org/zap"
)

// StraightSubtraction subtracts num2 from num1 then returns the result
func StraightSubtraction(num1, num2 int) int {
	return num1 - num2
}

// LogWithZap logs num with a zap logger
func LogWithZap(num int) {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Sync()
	logger.Info("Your num is", zap.Int("num", num))
}
