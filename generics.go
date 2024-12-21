package main

import (
	"fmt"
)

type Numeric interface {
	int | float64
}

func Sum[T Numeric](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(Sum(10, 20))
	fmt.Println(Sum(3.5, 4.5))
}
