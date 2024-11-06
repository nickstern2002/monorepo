package firstModule

import (
	"fmt"
	"github.com/mazen160/go-random"
)

func straightAddition(num1, num2 int) int {
	return num1 + num2
}

func randomNum() string {
	string, err := random.String(3)
	if err != nil {
		fmt.Println(err)
	}
	return string
}
