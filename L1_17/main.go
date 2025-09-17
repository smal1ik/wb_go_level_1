package main

import "fmt"

func binarySearch(a []int, number, leftIdx, rightIdx int) int {
	if rightIdx >= leftIdx {
		middleIdx := leftIdx + (rightIdx-leftIdx)/2
		if a[middleIdx] == number {
			return middleIdx
		} else if a[middleIdx] > number {
			return binarySearch(a, number, leftIdx, middleIdx-1)
		} else {
			return binarySearch(a, number, middleIdx+1, rightIdx)
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(a, 9, 0, len(a)-1))
	fmt.Println(binarySearch(a, 0, 0, len(a)-1))
}
