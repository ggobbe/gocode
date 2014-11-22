package main

func problem2() int {
	var (
		first  int
		second = 1
		next   int
		sum    int
	)

	for next < 4000000 {
		next = first + second
		first = second
		second = next

		if next%2 == 0 {
			sum += next
		}
	}
	return sum
}
