package main

func problem5() int {
	return smallestMultiple(20)
}

func smallestMultiple(dividendMax int) int {
	for i := dividendMax; ; i += dividendMax {
		if canBeDivided(i, dividendMax) {
			return i
		}
	}
}

func canBeDivided(n int, dividendMax int) bool {
	for i := dividendMax - 1; i > 0; i-- {
		if n%i != 0 {
			return false
		}
	}
	return true
}
