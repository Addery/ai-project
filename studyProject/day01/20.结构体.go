package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintInfo() {
	fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
}

func main() {
	s := Student{
		Name: "zzy",
		Age:  20,
	}
	s.Name = "Addery"
	s.PrintInfo()
}
