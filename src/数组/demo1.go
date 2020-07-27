package main

import "fmt"

func main() {
	// var 数组变量名 [元素数量]Type
	// 和切片区别就是没有指定长度
	var a [10]int
	var c [10]int
	fmt.Println(a, c)
	mylist := [1]int{1}
	mylist[0] = 10

	// 切片
	var myslice []int
	b := []int{1, 2}
	fmt.Println(myslice, b)

	mysl := make([]int, 10, 10)
	fmt.Println(mysl)

}
