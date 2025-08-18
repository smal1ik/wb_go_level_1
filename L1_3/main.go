package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var ch = make(chan int, 100)

func worker(id int) {
	for number := range ch {
		fmt.Println("worker ", id, " got ", number)
	}
}

func main() {
	workerCount, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < workerCount; i++ {
		go worker(i)
	}
	for {
		ch <- rand.Intn(100)
		time.Sleep(500 * time.Millisecond)
	}
}
