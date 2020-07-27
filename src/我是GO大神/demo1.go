package main

import "fmt"

func godemo() {
	fmt.Println(1111)
}
func main() {
	go godemo()
}
