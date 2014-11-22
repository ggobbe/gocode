package main

func problem12() int {
	var sum int
	factors := []int{}
	largest := 0

	for i := 1; len(factors) <= 500; i++ {
		sum += i

		factors = findFactors(sum)

		if len(factors) > largest {
			largest = len(factors)
		}
	}

	return sum
}

func findFactors(n int) []int {
	factors := []int{}

	for i := 1; i < n; i++ {
		if n%i == 0 {
			if n/i <= i {
				break
			}
			factors = append(factors, i, n/i)
		}
	}

	return factors
}
