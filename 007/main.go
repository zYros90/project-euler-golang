package main

import (
	"fmt"
	"new/utils"
)

func main() {
	var counter int
	for i := 2; true; i++ {
		if utils.IsPrime(i) {
			counter++
			if counter == 10001 {
				fmt.Println(i)
				break
			}
		}
	}
}
