package main

func problem9() int {
	for a := 1; a <= 997; a++ {
		for b := a + 1; b <= 997; b++ {
			for c := b + 1; c <= 997; c++ {
				if a+b+c == 1000 && (a*a)+(b*b) == (c*c) {
					return a * b * c
				}
			}
		}
	}
	return 0
}
