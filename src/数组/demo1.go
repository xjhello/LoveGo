package main

import "fmt"

func main() {
	// var 数组变量名 [元素数量]Type
	// 和切片区别就是没有指定长度
	mylist := [1]int{1}
	mylist[0] = 10

	var myslice []int
	fmt.Println(myslice)

	mysl := make([]int, 10, 10)

}
