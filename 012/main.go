package main

import (
	"fmt"
)

func main() {
	for i := 2; true; i++ {
		t := getTriangularNumber(i)
		numDivisors := 2
		var divisorOld int
		for j := 2; j < t/2; j++ {
			if t%j != 0 {
				continue
			}
			divisor := t / j
			if divisor == j {
				numDivisors++
				break
			}
			if divisorOld == j {
				break
			}
			numDivisors += 2
			divisorOld = divisor
		}
		if numDivisors > 500 {
			fmt.Println("triangular number:", t)
			break
		}
	}
}

func getTriangularNumber(i int) int {
	return i * (i + 1) / 2
}
