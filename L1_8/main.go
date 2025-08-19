package main

import "fmt"

func setBit(num int64, i uint, bit int) int64 {
	if bit == 0 {
		return num & ^(1 << (i - 1))
	}
	return num | (1 << (i - 1))
}

func main() {
	var num int64 = 5

	fmt.Println(setBit(num, 1, 0))

	fmt.Println(setBit(num, 2, 1))
}
