package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func workerCondition() {
	fmt.Println("workerCondition run")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("workerCondition done")
			return
		}
	}
}

func workerChan(done <-chan bool) {
	fmt.Println("workerChan run")
	for {
		select {
		case <-done:
			fmt.Println("workerChan done")
			return
		default:
		}
	}
}

func workerContext(ctx context.Context) {
	fmt.Println("workerContext run")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("workerContext done")
			return
		default:
		}
	}
}

func workerGoexit() {
	fmt.Println("workerGoexit run")
	time.Sleep(1 * time.Second)
	fmt.Println("workerGoexit done")
	runtime.Goexit()
}

func workerPanic() {
	fmt.Println("workerPanic run")
	time.Sleep(1 * time.Second)
	panic("workerPanic done")
}

func main() {
	// Завершение горутины по условию
	go workerCondition()

	// Завершение горутины через канал
	stop := make(chan bool)
	go workerChan(stop)
	time.Sleep(time.Second * 1)
	stop <- true

	// Завершение горутины через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go workerContext(ctx)
	time.Sleep(time.Second * 1)
	cancel()

	// Завершение горутины через runtime.Goexit
	go workerGoexit()
	time.Sleep(time.Second * 1)

	// Завершение горутины через panic
	go workerPanic()
	time.Sleep(time.Second * 1)
}
