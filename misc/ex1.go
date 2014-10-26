package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{0, 3, 3, 7, 5, 3, 11, 1}
	fmt.Printf("Adjacent pairs: %d\n", numberOfAdjacentPairs(array))
}

type Node struct {
	Value int
	Low   int
	High  int
}

func numberOfAdjacentPairs(a []int) int {
	num := 0

	adj := make(map[int]*Node)

	for i := 0; i < len(a); i++ {
		if _, ok := adj[a[i]]; !ok {
			adj[a[i]] = &Node{a[i], a[i], a[i]}
		}

		for j := i + 1; j < len(a); j++ {
			if _, ok := adj[a[j]]; !ok {
				adj[a[j]] = &Node{a[j], a[j], a[j]}
			}

			if a[i] > a[j] && (adj[a[i]].Low < a[j] || adj[a[i]].Low == a[i]) {
				adj[a[i]].Low = a[j]
				adj[a[j]].High = a[i]
			}

			if a[i] < a[j] && (adj[a[i]].High > a[j] || adj[a[i]].High == a[i]) {
				adj[a[i]].High = a[j]
				adj[a[j]].Low = a[i]
			}

			if math.Abs(float64(a[i]-a[j])) > 1 {
				num++
			} else {
				num++
				fmt.Printf("Adjacent pair: a[%d]=%d - a[%d]=%d\n", i, a[i], j, a[j])
			}
		}
	}

	for key, value := range adj {
		fmt.Printf("[%d]Node %d [%d, %d]\n", key, value.Value, value.Low, value.High)
	}

	return num
}
