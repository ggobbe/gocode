package main

func problem6() int {
	return sumSquareDiff(100)
}

func sumSquareDiff(n int) int {
	var sumSquared int
	for i := 1; i <= n; i++ {
		sumSquared += i * i
	}

	sumAll := (n * (n + 1)) / 2

	return sumAll*sumAll - sumSquared
}
