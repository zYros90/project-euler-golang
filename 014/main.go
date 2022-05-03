package main

import "fmt"

func main() {
	start := 1000000
	maxLengthChain := 0
	maxStartingNr := 0
	for i := start; i > 0; i-- {
		x := i
		number := 1
		for x > 1 {
			if x%2 == 0 {
				x = x / 2
				number++
			} else {
				x = 3*x + 1
				number++
			}
		}
		number++
		if number > maxLengthChain {
			maxLengthChain = number
			maxStartingNr = i
		}
	}
	fmt.Println("starting number:", maxStartingNr)
}
