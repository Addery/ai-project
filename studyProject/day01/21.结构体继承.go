package main

import "fmt"

type People struct {
	Time string
}

func (p People) Info() {
	fmt.Println("people: ", p.Time)
}

type Teacher struct {
	People
	Name string
	Age  int
}

func (t Teacher) PrintInfo() {
	fmt.Printf("name: %s age: %d\n", t.Name, t.Age)
}

func main() {
	p := People{
		Time: "2026-2-4 20:58",
	}

	t := Teacher{
		People: p,
		Name:   "zzy",
		Age:    18,
	}
	t.Name = "Addery"

	t.Info()
	t.PrintInfo()

	fmt.Println(t.People.Time)
	fmt.Println(t.Name)
}
