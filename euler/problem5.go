package main

import "fmt"

func main() {
	fmt.Println("Answer:", smallestMultiple(20))
}

func smallestMultiple(dividendMax int) int {
	for i := 1; ; i++ {
		if canBeDivided(i, dividendMax) {
			return i
		}
	}
}

func canBeDivided(n int, dividendMax int) bool {
	for i := 1; i <= dividendMax; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}
