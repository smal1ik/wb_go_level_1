package main

import (
	"fmt"
)

func getTypeVar(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("var type int")
	case string:
		fmt.Println("var type string")
	case bool:
		fmt.Println("var type bool")
	case chan bool:
		fmt.Println("var type chan bool")
	case chan int:
		fmt.Println("var type chan int")
	case chan string:
		fmt.Println("var type chan string")
	}

}

func main() {
	a := 1
	getTypeVar(a)

	b := true
	getTypeVar(b)

	c := "hello"
	getTypeVar(c)

	d := make(chan int)
	getTypeVar(d)

	e := make(chan bool)
	getTypeVar(e)

	f := make(chan string)
	getTypeVar(f)
}
