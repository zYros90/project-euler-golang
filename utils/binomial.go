package utils

import (
	"math/big"
)

func GetBinomialCoefficient(n int, k int) int64 {
	nFactorial := GetFactorial(n)
	kFactorial := GetFactorial(k)
	nMinkFactorial := GetFactorial(n - k)
	quotient := kFactorial.Mul(kFactorial, nMinkFactorial)
	result := nFactorial.Div(nFactorial, quotient)
	// return GetFactorial(n) / (GetFactorial(k) * GetFactorial(n-k))
	return result.Int64()
}

func GetFactorial(n int) *big.Int {
	factorial := big.NewInt(1)
	for i := int64(1); i <= int64(n); i++ {

		factorial = factorial.Mul(factorial, big.NewInt(i))
	}
	if factorial == big.NewInt(0) {
		return big.NewInt(1)
	}
	return factorial
}
