package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var ch = make(chan int, 100)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker ", id, " done")
			return
		case number := <-ch:
			fmt.Println("worker ", id, " got ", number)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		<-sigCh
		cancel()
	}()

	workerCount, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < workerCount; i++ {
		go worker(ctx, i)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case ch <- rand.Intn(100):
		}
		time.Sleep(500 * time.Millisecond)
	}
}
