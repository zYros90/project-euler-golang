package main

import "fmt"

func main() {
	n := 10000
	sum := 0

	for i := 2; i < n; i++ {
		divisorSum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				divisorSum += j
			}
		}

		if isAmicableNumber(i, divisorSum) {
			sum += i
		}
	}

	fmt.Println("sum: ", sum)
}

func isAmicableNumber(origin int, divisorSum int) bool {
	newDivisorSum := 0
	for i := 1; i < divisorSum; i++ {
		if divisorSum%i == 0 {
			newDivisorSum += i
		}
	}

	return newDivisorSum == origin && divisorSum != newDivisorSum
}
