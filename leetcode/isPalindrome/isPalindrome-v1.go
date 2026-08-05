package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := numSize(x)

	p := s % 2 == 0

	var m int = -1
	if p {
		m = s / 2
	}

	for i := 0; x >= 10; i++ {
		r := x % 10

		// Skip the middle number
		if i == m {
			x = x / 10
			continue
		}

		// if r is different from the top number then false
		if r != x / int(math.Pow(10, float64(s - i))) {
			return false
		}
		// x looses top number
		x = x - int(math.Pow(10, float64(s)))
		// size is one less
		s = s - 1
		// get rid of last number
		x = x / 10
	}
	
	return true
}

func numSize(x int) int {
	var i int
	for x >= 10 {
		x = x / 10
		i++
	}

	return i
}

func main() {
	fmt.Println(isPalindrome(2))
}
