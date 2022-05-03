package utils

// n needs to be n > 1
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// get prime factors with sieve of Erathosthenes
// achieves O(log n) for all composite numbers but is O(n) otherwise
func SieveOfEratosthenes(n int) []int {
	primeFactors := make([]int, 0)
	c := 2
	for n > 1 {
		if n%c == 0 {
			primeFactors = append(primeFactors, c)
			n /= c
			continue
		}
		c += 1
	}
	return primeFactors
}
