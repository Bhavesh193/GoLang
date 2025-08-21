package main

import (
	"fmt"
	"time"
)

var counter = 0

func increment() {
	for i := 0; i < 5; i++ {
		counter++
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go increment()
	go increment()

	time.Sleep(time.Second) // wait for goroutines to finish
	fmt.Println("Final Counter:", counter)
}
