package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// this goroutine waits for a value from the channel, but it will never receive one because the second goroutine never sends one
	go func() {
		val := <-ch
		fmt.Println(val)
	}()

	// this goroutine sends a value to the channel and then waits for the first goroutine to complete, but it will never complete because the first goroutine is waiting for a value that will never be sent
	go func() {
		ch <- 42
		<-ch
	}()

	// wait for both goroutines to finish before exiting
	select {}
}
