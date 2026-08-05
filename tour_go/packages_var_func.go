package main

import (
	"fmt"
	"math"
)

// If both share a type (x, y int) == (x int, y int)
func add(x int, y int) int {
	return x + y
}

// You can have multiple returns (Takes two strings and returns two strings)
func swap(x, y string) (string, string) {
	return y, x
}
/*
 * Go's return values may be named. If so, they are treated as variables defined at the top of the function.
 * These names should be used to document the meaning of the return values.
 * A return statement without arguments returns the named return values. This is known as a "naked" return.
 * Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions. 
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variable declaration works as with function arguments
var a, b int

// Can initialize while declaring, also works with multiple vars
var i, j int = 1, 2

// Type can be omitted
var c, python, java = true, false, "no!"

func main()  {
	a, b := swap("hello", "world")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(add(42, 43))
}
