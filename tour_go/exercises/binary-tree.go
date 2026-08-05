package main

import (
	"golang.org/x/tour/tree"
	"fmt"
) 

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	func() { Walk(t1, ch1); close(ch1) } ()
	func() { Walk(t2, ch2); close(ch2) } ()

	for v1 := range ch1 {
		if v2, ok := <- ch2; v2 != v1 || !ok {
			return false
		}
	}
	_, ok := <- ch2

	return !ok
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
