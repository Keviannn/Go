package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Ifs are the same in the sense that () are not needed but {} do
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Ifs can have its own declarations before the evaluation, the scope lasts until the if dies
// If there is an else the scope lasts until then
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// Switch case example
/*
 * In effect, the break statement that is needed at the end of each case in those 
 * languages is provided automatically in Go. 
 *
 * Another important difference is that Go's switch cases need not be constants, 
 * and the values involved need not be integers
 */
// Switch cases evaluate cases from top to bottom, stopping when a case succeeds. 
func os() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	// switcht without evaluation is a clean if-then-else statement
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func test() {
	sum := 0
	// Go only has for as a loop, () not needed but {} do
	for i := 0; i < 10; i++ {
		sum += i
	}

	// Can be simplified
	for ; sum < 1000; {
		sum += sum
	}

	// So much you can make a while loo getting rid of the ;
	for sum < 1000 {
		sum += sum
	}

	// Infinite loops can be like
	/*
	 * for {

	 * }
	 */

	// Defer waits for the sorrounding function (main) to end before executing
	// Its arguments are evaluated at the normal time
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")

	// Those calls are put into a stack LIFO
	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
