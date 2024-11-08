package firstModule

import (
	"fmt"
	"github.com/mazen160/go-random"
)

// StraightAddition adds two numbers
func StraightAddition(num1, num2 int) int {
	return num1 + num2
}

// RandomNum actually returns a string lol
func RandomNum() string {
	string, err := random.String(3)
	if err != nil {
		fmt.Println(err)
	}
	return string
}
