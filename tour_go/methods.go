package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Also for basic types
type MyFloat float64

// You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer recievers make methods able to edit values, as value recievers only work on copies
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// They take v.Scale and (&v).Scale the same way, so you can pass pointers or values as Go interprets the first 
// as the last, but is essentially working as the last
// The same is true for value recievers, they can be called with p = &v, p.call() as they will be interpreted as
// (*p).call()

// Choosing pointer over values is because modification and avoiding copies,
// Mixing types of methods for the same type isn't a good practice

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v.Abs())
}
