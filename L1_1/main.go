package main

import "fmt"

type Human struct {
	Age        int
	FirstName  string
	SecondName string
	Sex        string
}

type Action struct {
	Human
}

func (h Human) GetFullName() string {
	return h.FirstName + " " + h.SecondName
}

func (h Human) GetYearBirthday() int {
	return 2025 - h.Age
}

func (a Action) Walk() {
	fmt.Println(a.GetFullName(), "walking...")
}

func (a Action) Jump() {
	fmt.Println(a.FirstName, "jumping...")
}

func main() {
	human := Human{
		Age:        18,
		FirstName:  "Test",
		SecondName: "Testov",
		Sex:        "male",
	}

	action := Action{
		Human: human,
	}

	action.Walk()
	action.Jump()

	fmt.Println(action.GetFullName())
	fmt.Println(action.GetYearBirthday())
}
