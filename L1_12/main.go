package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, word := range words {
		_, ok := set[word]
		if !ok {
			set[word] = struct{}{}
		}
	}

	result := []string{}

	for key := range set {
		result = append(result, key)
	}
	fmt.Println(result)

}
