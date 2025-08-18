package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch = make(chan int, 100)

func worker() {
	for number := range ch {
		fmt.Printf("get %d \n", number)
	}
}

func main() {
	go worker()

	timeout := time.After(3 * time.Second)

	for {
		select {
		case <-timeout:
			close(ch)
			return
		case ch <- rand.Intn(100):
		}
	}
}
