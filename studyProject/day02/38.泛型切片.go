package main

type MySlice[T any] []T

func main() {
	var mySlice MySlice[string]
	mySlice = append(mySlice, "枫枫")
	var intSlice MySlice[int]
	intSlice = append(intSlice, 2)
}
