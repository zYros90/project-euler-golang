package main

import "fmt"

func main() {
	for i := 1; true; i++ {
		foundSmallestNumber := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				foundSmallestNumber = false
			}
		}
		if foundSmallestNumber {
			fmt.Println("smallest number: ", i)
			break
		}
	}
}
