package main

import (
	"fmt"
	"new/utils"
)

// pascals triangle
// https://en.wikipedia.org/wiki/Pascal%27s_triangle
// n: sum of edges top and right => n = 2xdim
// k = n/2
// for 2x2 grid, n=2x2, k=2 => 6
// for 20x20 grix, n=20x20, k=20
func main() {
	fmt.Println(utils.GetBinomialCoefficient(40, 20))
}
