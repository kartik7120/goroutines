package main

import (
	"fmt"
	"time"
)

func someFunc() {
	fmt.Println("This is someFunc")
}

func main() {
	go someFunc()
	go someFunc()
	go someFunc()

	time.Sleep(time.Second * 2)

	fmt.Println("Hello, World!")
}

/*
	Go follows a model of concurrency called Fork-Join model.
	In order to communicate between goroutines, we can use channels.
*/
