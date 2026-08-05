package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	var sum3 int
	return func() int {
		sum3 = sum1 + sum2
		sum1 = sum2
		sum2 = sum3
		return sum3
	}	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
