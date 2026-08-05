package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := numSize(x)	

	v := make([]int, s + 1)
	for i := 0; x >= 10; i++ {
		v[i] = x % 10
		x = x / 10
	}
	v[s] = x

	for i := range v {
		if v[i] != v[s - i] {
			return false
		}
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
	fmt.Printf("It is %v", isPalindrome(1221))
}
