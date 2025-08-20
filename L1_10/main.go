package main

import "fmt"

func main() {
	var a = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	ran := make(map[int][]float32)
	var decimal int
	for _, num := range a {
		decimal = int(num/10) * 10
		ran[decimal] = append(ran[decimal], num)
	}
	fmt.Println(ran)
}
