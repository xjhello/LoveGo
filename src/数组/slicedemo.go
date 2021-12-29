package main

import "fmt"


func main(){
	a := []string{"ssss", "sdasda"}

	// b := make([]int, 0)
	fmt.Println(a)

	a = append(a, "@@@@")
	fmt.Println(a)
	for i:=0; i<len(a); i++ {
		fmt.Println(i)
	}

	for index, value := range a {
		fmt.Println(index, value)
	}
	
	aa := "asdasd"
	bb := fmt.Sprintf(">>> %s", aa)
	print(bb)
	// for index, value := range s {
	// 	fmt.Println(index, value)
	// }
}