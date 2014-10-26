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
		for j := i + 1; j < len(a); j++ {
			if math.Abs(float64(a[i]-a[j])) > 1 {
				num++
			} else {
				if _, ok := adj[a[i]]; !ok {
					adj[a[i]] = &Node{a[i], a[i], a[i]}
				}
				if _, ok := adj[a[j]]; !ok {
					adj[a[j]] = &Node{a[j], a[j], a[j]}
				}
				num++
				fmt.Printf("Adjacent pair: a[%d]=%d - a[%d]=%d\n", i, a[i], j, a[j])
			}
		}
	}

	return num
}
