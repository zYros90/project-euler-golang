package main

import (
	"fmt"
	"math"
)

func main() {
	var sumSquares float64
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		sumSquares += math.Pow(float64(i), 2.0)
	}
	fmt.Println("diff: ", math.Pow(float64(sum), 2)-sumSquares)
}
