package secondModul

import (
	"fmt"
	"go.uber.org/zap"
)

func straightSubtraction(num1, num2 int) int {
	return num1 - num2
}

func logWithZap(num int) {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Sync()
	logger.Info("Your num is", zap.Int("num", num))
}
