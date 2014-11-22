package main

import "math"

func problem10() int {
	return sumOfPrimes(2000000)
}

func sumOfPrimes(max int) int {
	primes := []int{2}
	sum := 2

	for i := 3; i < max; i += 2 {
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
		sum += i
	}
	return sum
}
