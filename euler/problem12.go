package main

func problem12() int {
	var sum int
	factors := 0

	for i := 1; factors <= 500; i++ {
		sum += i

		factors = countFactors(sum)
	}

	return sum
}

func countFactors(n int) int {
	factors := 0

	for i := 1; n/i > i; i++ {
		if n%i == 0 {
			factors += 2
		}
	}

	return factors
}
