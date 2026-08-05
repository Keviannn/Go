package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64
	z = 1
	for range 10 {
		z -= (z*z -x) / (2*z)
		fmt.Println(z)
	}
	return float64(z)
}

func main() {
	fmt.Println(Sqrt(4))
}
