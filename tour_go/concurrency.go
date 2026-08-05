package main

import "fmt"

// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
 * Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
 * ch := make(chan int, 100)
 *
 * Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
 */

/*
 * A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
 * 
 * v, ok := <-ch
 * 
 * ok is false if there are no more values to receive and the channel is closed.
 * 
 * The loop for i := range c receives values from the channel repeatedly until it is closed.
 * 
 * Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
 * 
 * Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
 */

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// consumes from p until its closed
	p := make(chan int, 10)
	go fibonacci(cap(p), p)
	for i := range p {
		fmt.Println(i)
	}
}
