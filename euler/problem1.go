package main

func problem1() int {
	var sum int
	for i := 3; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	return sum
}
