package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < len(a); i++ {
		if a[i] < pivot {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)

}

func main() {
	a := []int{5, 4, 3, 2, 1, 10, 22, 9, 6, 6, 50, 0, 3, 5}
	fmt.Println(quickSort(a))
}
