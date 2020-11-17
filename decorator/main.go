package main

import "fmt"

func main() {
	var window = NewVerticalScrollBars(
		NewVerticalScrollBars(new(SimpleWindow)))
	fmt.Println(window.GetDescription())

	window = NewVerticalScrollBars(
		NewVerticalScrollBars(new(ComplexWindow)))
	fmt.Println(window.GetDescription())

}
