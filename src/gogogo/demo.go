package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a, len(a[1:3]), a[1:3])
	mylist := make([]int, 10)
	mylist[9] = 10
	println(mylist)

	mmp := map[string]int{"1": 1}
	mmmp := make(map[string]string)
	mmmp["2"] = "sad"
	mmmp["1"] = "sad"
	fmt.Println(mmp, mmmp)
	test(11)
}

func test(a int) {
	fmt.Println(a)
}
