package main

import "fmt"

// Structs
type Vertex struct {
	X int
	Y int
}
// Pointers to structcs can access values with p.X avoiding (*p).X
// A struct literal denotes a newly allocated struct value by listing the values of its fields.
// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
// The special prefix & returns a pointer to the struct value.

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {

	// Pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Strings
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array
	// Primes 1 to <4 so positions 1,2,3
	var s []int = primes[1:4]
	fmt.Println(s)


	// Slices does not store data
	// Changing it changes the array and therefore other slices
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	m := names[0:2]
	k := names[1:3]
	fmt.Println(a, b)

	k[0] = "XXX"
	fmt.Println(m, k)
	fmt.Println(names)


//A slice literal is like an array literal without the length.

// This is an array literal:
// [3]bool{true, true, false}

// And this creates the same array as above, then builds a slice that references it:
// []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	l := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(l)
/*
 * When slicing, you may omit the high or low bounds to use their defaults instead.
 * 
 * The default is zero for the low bound and the length of the underlying slice or array for the high bound.
 * 
 * For the array
 * 
 * var a [10]int
 * 
 * these slice expressions are equivalent:
 * 
 * a[0:10]
 * a[:10]
 * a[0:]
 * a[:]
 */

// Un slice tiene len(s) y cap(s), len es su longitud real, cap es la cantidad de elementos desde el incio
// del slice al final del array que referencia

// a nil slice is a slice with no elements and len and cap are zero

/*
 * Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
 * 
 * The make function allocates a zeroed array and returns a slice that refers to that array:
 * 
 * a := make([]int, 5)  // len(a)=5
 * 
 * To specify a capacity, pass a third argument to make:
 * 
 * b := make([]int, 0, 5) // len(b)=0, cap(b)=5
 * 
 * b = b[:cap(b)] // len(b)=5, cap(b)=5
 * b = b[1:]      // len(b)=4, cap(b)=4
 */

// Pueden tener slices de slices
/*
 * board := [][]string{
 * 		[]string{"_", "_", "_"},
 * 		[]string{"_", "_", "_"},
 * 		[]string{"_", "_", "_"},
 * 	}
 * 
 * 	// The players take turns.
 * 	board[0][0] = "X"
 * 	board[2][2] = "O"
 * 	board[1][2] = "X"
 * 	board[1][0] = "O"
 * 	board[0][2] = "X"
 */

// Append
/*
 *  append works on nil slices.
 * 	s = append(s, 0)
 * 	printSlice(s)
 * 
 * 	// The slice grows as needed.
 * 	s = append(s, 1)
 * 	printSlice(s)
 * 
 * 	// We can add more than one element at a time.
 * 	s = append(s, 2, 3, 4)
 * 	printSlice(s)
 */
// If the backing array of s is too small to fit all the given values a bigger array will be allocated. 
// The returned slice will point to the newly allocated array. 

// Range sirve para recorrer un slice
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
/*
 * 2**0 = 1
 * 2**1 = 2
 * 2**2 = 4
 * 2**3 = 8
 */

/* 
 *   You can skip the index or value by assigning to _.
 * 
 * for i, _ := range pow
 * for _, value := range pow
 * 
 * If you only want the index, you can omit the second variable.
 * 
 * for i := range pow
 */


}

