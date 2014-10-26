// Cycle detection: Floyd's cycle-finding algorithm

package main

import (
	"fmt"
)

func main() {
	array := []int{6, 6, 0, 1, 4, 3, 3, 4, 0}

	lam, mu := floyd(array)
	fmt.Printf("λ=%d, μ=%d\n", lam, mu)
}

func floyd(a []int) (int, int) {
	tortoise := a[0]
	hare := a[a[0]]

	// find the repetition
	for tortoise != hare {
		tortoise = a[tortoise]
		hare = a[a[hare]]
	}

	mu := 0
	tortoise = 0

	// find the position µ of the first repetition
	for tortoise != hare {
		tortoise = a[tortoise]
		hare = a[hare]
		mu++
	}

	lam := 1
	hare = a[tortoise]

	// find the length of the shortest cycle
	for tortoise != hare {
		hare = a[hare]
		lam++
	}

	return lam, mu
}
