package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	e := ErrNegativeSqrt(x)
	if x < 0 {
		//fmt.Printf("Tipo: %T, Valor: %v\n", e, e)
		return 0, e
	}

	var z float64
	z = 1
	for range 10 {
		z -= (z*z -x) / (2*z)
		fmt.Println(z)
	}
	return float64(z), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
