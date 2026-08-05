package main

import (
	"fmt"
	"math"
)

/*
 * A type implements an interface by implementing its methods. 
 * There is no explicit declaration of intent, no "implements" keyword.
 * 
 * Implicit interfaces decouple the definition of an interface from its implementation, 
 * which could then appear in any package without prearrangement.
 * 
 * Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
 * 
 * (value, type)
 * 
 * An interface value holds a value of a specific underlying concrete type.
 * 
 * Calling a method on an interface value executes the method of the same name on its underlying type.
 */

/*
 * If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
 * 
 * In some languages this would trigger a null pointer exception, but in 
 * Go it is common to write methods that gracefully handle being called with a 
 * nil receiver (as with the method M in this example.)
 * 
 * Note that an interface value that holds a nil concrete value is itself non-nil.
 */

/*
 * For example when creating a var i I form an I interface without definition only declaration.
 * A nil interface value holds neither value nor concrete type.
 * 
 * Calling a method on a nil interface is a run-time error because there is no type 
 * inside the interface tuple to indicate which concrete method to call.
 */

/*
 * The interface type that specifies zero methods is known as the empty interface:
 * interface{}
 * 
 * An empty interface may hold values of any type. (Every type implements at least zero methods.)
 * 
 * Empty interfaces are used by code that handles values of unknown type. 
 * For example, fmt.Print takes any number of arguments of type interface{}.

 func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

(<nil>, <nil>)
(42, int)
(hello, string)

*/

/*
 * A type assertion provides access to an interface value's underlying concrete value.
 * 
 * t := i.(T)
 * 
 * This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
 * 
 * If i does not hold a T, the statement will trigger a panic.
 * 
 * To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
 * 
 * t, ok := i.(T)
 * 
 * If i holds a T, then t will be the underlying value and ok will be true.
 * 
 * If not, ok will be false and t will be the zero value of type T, and no panic occurs.
 * 
 * Note the similarity between this syntax and that of reading from a map.

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
	/tmp/sandbox1792448541/prog.go:17 +0x13f
*/

// Switch can check type of an interface
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	// a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
