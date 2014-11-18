package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Answer:", largestPalindromeProduct(3))
}

func largestPalindromeProduct(n int) int {
	var largest int

	start := int(math.Pow10(n - 1))
	end := int(math.Pow10(n))

	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			product := i * j
			if product > largest && isPalindrome(product) {
				largest = product
			}
		}
	}
	return largest
}

func isPalindrome(n int) bool {
	return n == reverse(n)
}

func reverse(n int) int {
	var reverse int

	for n != 0 {
		reverse *= 10
		reverse += n % 10
		n /= 10
	}
	return reverse
}
