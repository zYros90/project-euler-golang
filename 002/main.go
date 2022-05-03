package main

import "fmt"

func main() {
	prev := 1
	new := 2
	sum := 0
	for i := new; new < 4000000; i++ {
		if new%2 == 0 {
			sum += new
		}
		prev, new = new, prev+new
	}
	fmt.Println("sum: ", sum)
}
