package main

import "fmt"

var ch1 = make(chan int, 100)
var ch2 = make(chan int, 100)

var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func workerFirst() {
	for _, elem := range array {
		ch1 <- elem
	}
	close(ch1)
}

func workerSecond() {
	for number := range ch1 {
		ch2 <- number * 2
	}
	close(ch2)
}

func main() {

	go workerFirst()
	go workerSecond()

	for number := range ch2 {
		fmt.Println(number)
	}
}
