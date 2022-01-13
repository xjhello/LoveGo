package main

import "fmt"

func returnDemo(a int) *int {
	return &a
}

func main() {
	x := 10
	fmt.Println(returnDemo(x))
	ptr := returnDemo(x)
	fmt.Println(ptr)
}
