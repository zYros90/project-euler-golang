package main

import (
	"fmt"
)

func main() {
	n := 600851475143
	fmt.Println("primefactors: ", SieveOfEratosthenes(n))
}

func SieveOfEratosthenes(n int) []int {
	primeFactors := make([]int, 0)
	c := 2
	for n > 1 {
		if n%c == 0 {
			primeFactors = append(primeFactors, c)
			n /= c
			continue
		}
		c += 1
	}
	return primeFactors
}
