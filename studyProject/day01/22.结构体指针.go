package main

import "fmt"

type Worker struct {
	Name string
	Age  int
}

func setAge(info Worker, age int) {
	info.Age = age
}

func setAge1(info *Worker, age int) {
	info.Age = age
}

func (w Worker) setAge2(age int) {
	w.Age = age
}

func (w *Worker) setAge3(age int) {
	w.Age = age
}

func main() {
	w := Worker{
		Name: "Addery",
		Age:  20,
	}
	fmt.Println(w.Age)
	setAge(w, 18)
	fmt.Println(w.Age)
	setAge1(&w, 16)
	fmt.Println(w.Age)
	
	w.setAge2(14)
	fmt.Println(w.Age)
	w.setAge3(12)
	fmt.Println(w.Age)

}
