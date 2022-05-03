package main

import (
	"fmt"
	"new/utils"
)

func main() {
	var sum int
	for i := 2; i < 2000000; i++ {
		if utils.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println("sum:", sum)
}
