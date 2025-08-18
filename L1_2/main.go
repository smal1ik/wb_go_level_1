package main

import (
	"fmt"
	"sync"
)

var array = []int{2, 4, 6, 8, 10}
var wg sync.WaitGroup

func main() {
	for _, elem := range array {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(elem)
	}
	wg.Wait()
}
