package main

import (
	"math/rand"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (s *SafeMap) Get(key int) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SafeMap) Set(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	sm := SafeMap{m: make(map[int]int)}

	for i := 0; i < 100; i++ {
		go func() {
			sm.Set(rand.Intn(100), rand.Intn(100))
		}()
	}
}
