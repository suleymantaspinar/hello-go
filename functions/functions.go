package functions

import (
	"fmt"
)

func Run(from int, to int) {
	for i := 0; i < from; i++ {
		fmt.Print(fmt.Sprintf("%d-", i))
	}
	for i := 0; i < to; i++ {
		fmt.Print(" to", i)

	}
}

func Multiply(a, b int) int {
	return a + b
}
