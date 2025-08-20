package main

import "fmt"

func main() {
	var A = []int{1, 2, 3}
	var B = []int{2, 3, 4}
	var hash = make(map[int]struct{})
	var intersection []int

	for _, elem := range A {
		hash[elem] = struct{}{}
	}

	for _, elem := range B {
		_, ok := hash[elem]
		if ok {
			intersection = append(intersection, elem)
		}

	}

	fmt.Println(intersection)

}
