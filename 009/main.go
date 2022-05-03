package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			cFloat64 := math.Sqrt(float64(a*a + b*b))
			if isNaturalNumber(cFloat64) {
				if isSolution(a, b, int(cFloat64)) {
					fmt.Println(a * b * int(cFloat64))
					return
				}
			}
		}
	}
}

func isSolution(a, b, c int) bool {
	return a+b+c == 1000
}

func isNaturalNumber(c float64) bool {
	return c == float64(int64(c))
}
