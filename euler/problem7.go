package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Answer:", primeNth(10001))
}

func primeNth(n int) int {
	if n == 1 {
		return 2
	}

	primes := []int{2}
	count := 1

	for i := 3; ; i += 2 {
		var factorFound bool

		for _, p := range primes {
			if i%p == 0 {
				factorFound = true
				break
			}

			if float64(p) > math.Sqrt(float64(i)) {
				break
			}
		}

		if factorFound {
			continue
		}

		primes = append(primes, i)
		count++
		if count == n {
			return i
		}
	}
}
