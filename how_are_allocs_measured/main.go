package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")
}

//go:noinline
func OnTheHeap() *int {
	res := _onTheHeap()
	*res *= 10
	return res
}

//go:noinline
func _onTheHeap() *int {
	y := 2
	res := y * 2
	return &res
}

func OnTheStack() int {
	res := _onTheStack()
	res *= 10
	return res
}
func _onTheStack() int {
	y := 2
	return y * 2
}
