package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	n := int64(100)
	faculty := big.NewInt(1)

	for i := n; i > 0; i-- {
		faculty = faculty.Mul(faculty, big.NewInt(i))
	}

	sum := 0
	for _, x := range faculty.String() {
		xNum, err := strconv.Atoi(string(x))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += xNum
	}

	fmt.Println("sum: ", sum)
}
