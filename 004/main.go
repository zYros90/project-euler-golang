package main

import "fmt"

func main() {
	maxPalindrome := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			x := i * j
			if isPalindrome(x) {
				if x > maxPalindrome {
					maxPalindrome = x
				}
			}
		}
	}
	fmt.Println("max palindrome: ", maxPalindrome)
}

func isPalindrome(x int) bool {
	r := []rune(fmt.Sprintf("%d", x))
	if len(r)%2 != 0 {
		return false
	}
	rightSide := r[len(r)/2:]
	leftSide := make([]rune, 0)
	for i := len(r)/2 - 1; i >= 0; i-- {
		leftSide = append(leftSide, r[i])
	}
	return fmt.Sprintf("%v", rightSide) == fmt.Sprintf("%v", leftSide)
}
