package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.count++
}

func worker(c *Counter, nStep int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < nStep; i++ {
		c.mu.Lock()
		c.Increment()
		c.mu.Unlock()
	}
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(&counter, 320, &wg)
	wg.Add(1)
	go worker(&counter, 150, &wg)
	wg.Add(1)
	go worker(&counter, 100, &wg)

	wg.Wait()
	fmt.Println(counter.count)
}
