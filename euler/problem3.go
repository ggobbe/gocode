package main

func problem3() int {
	return largestPrimeFactor(600851475143)
}

func largestPrimeFactor(n int) int {
	largest := 1
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
