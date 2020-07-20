package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			goto mytag
		}
	}
mytag:
	fmt.Println("HERE!")

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}
