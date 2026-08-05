package main

import (
	"fmt"
	"math"
)

type Vertex_2 struct {
	Lat, Long float64
}

// Maps keys to values 
// zero value is nil, nil maps dont have keys nor can be added
var m map[string]Vertex_2

// map literals
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// you can omit the Vertex as it is assumed
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

/*
 * Insert or update an element in map m:
 * m[key] = elem
 * 
 * Retrieve an element:
 * elem = m[key]
 * 
 * Delete an element:
 * delete(m, key)
 * 
 * Test that a key is present with a two-value assignment:
 * elem, ok = m[key]
 * 
 * If key is in m, ok is true. If not, ok is false.
 * If key is not in the map, then elem is the zero value for the map's element type.
 * 
 * Note: If elem or ok have not yet been declared you could use a short declaration form:
 * elem, ok := m[key]
 */


// Functions are values too, can be sent and returned
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
/*
 * Go functions may be closures. A closure is a function value that references variables from outside its body. 
 * The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
 * 
 * For example, the adder function returns a closure. Each closure is bound to its own sum variable.
 */

func main() {
	m = make(map[string]Vertex_2)
	m["Bell Labs"] = Vertex_2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
