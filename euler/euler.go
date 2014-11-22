package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	displayAnswer(1, problem1)
	displayAnswer(2, problem2)
	displayAnswer(3, problem3)
	displayAnswer(4, problem4)
	displayAnswer(5, problem5)
	displayAnswer(6, problem6)
	displayAnswer(7, problem7)
	displayAnswer(8, problem8)
	displayAnswer(9, problem9)
	displayAnswer(10, problem10)
	displayAnswer(11, problem11)
	displayAnswer(12, problem12)
	stop := time.Now()
	fmt.Println("\t\t\t\tTotal:", stop.Sub(start))
}

func displayAnswer(n int, fn func() int) {
	start := time.Now()
	result := fn()
	stop := time.Now()
	fmt.Printf("Answer %2d: %15d\t(%s)\n", n, result, stop.Sub(start))
}
