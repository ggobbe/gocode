package main

import "fmt"

func main() {
	fmt.Println("Answer:", largestPrimeFactor(600851475143))
}

func largestPrimeFactor(n int) int {
	var largest int = 1
	for i := 2; n > 1; i++ {
		for n%i == 0 {
			if i > largest {
				largest = i
			}
			n /= i
		}
	}
	return largest
}
