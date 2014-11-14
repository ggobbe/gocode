package main

import "fmt"

func main() {
	var (
		first  int = 0
		second int = 1
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
	fmt.Println("Answer:", sum)
}
