// Exercise 3

package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []int64{33, 51, 114, 48763, 84777363}
	for _, v := range numbers {
		fmt.Printf("%d: %d\n", v, numberOfSteps(strconv.FormatInt(v, 2)))
	}
	fmt.Printf("Big one: %d\n", numberOfSteps("001010101110101001001010101001"+
		"00010100111101010101111010010111010100101111001010101111010101101111"+
		"10010101001010100010010010101001010111101010111101010100001010110101"+
		"1001101001011010101000001010111110101010010100010001011111101000101"))
}

// takes a string that represent an integer in its binary form and return the
// number of steps required to get to 0 by dividing the number by two when it
// is even and by subtracting 1 when the number is odd
func numberOfSteps(s string) int {
	if len(s) > 10101 {
		return -1
	}
	steps := -1
	for _, c := range s {
		if steps >= 0 && c == '0' {
			steps += 1
		} else if c == '1' {
			steps += 2
		} else if c != '0' && c != '1' {
			return -1
		}
	}
	return steps
}
