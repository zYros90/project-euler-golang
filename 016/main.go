package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(2)
	e := big.NewInt(1000)
	n.Exp(n, e, nil)
	var sum int64
	for {
		new := big.NewInt(0)
		new.Set(n)

		new.Mod(new, big.NewInt(10))
		n.Div(n, big.NewInt(10))
		sum += new.Int64()
		if n.Cmp(big.NewInt(1)) < 0 {
			break
		}
	}
	fmt.Println("sum:", sum)
}
