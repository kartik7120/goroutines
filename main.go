package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch <- "Hello, World!"
	}()

	go func() {
		ch2 <- "Hello, World! 2"
	}()

	select {
	case msg1 := <-ch:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

/*
	Go follows a model of concurrency called Fork-Join model.
	In order to communicate between goroutines, we can use channels.
*/
