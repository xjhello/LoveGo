package main

import "fmt"
import "mypkg/"

func godemo() {
	fmt.Println(1111)
}

func main() {
	go godemo()
}
