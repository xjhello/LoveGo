package main

import "fmt"

func main() {
	// var 数组变量名 [元素数量]Type
	// 和切片区别就是没有指定长度
	var a [10]int
	var c [10]int
	var ll [3]int
	ll[1] = 10
	fmt.Println(a, c)

	// 切片
	var myslice []int
	fmt.Println(myslice)
	b := []int{1, 2}
	fmt.Println(myslice, b)
	// 定义方式二
	mysl := make([]int, 10, 10)
	fmt.Println(mysl)

	a1 := []int{1, 2, 3, 4, 5, 7}
	s1 := a1[:]
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)

	//a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	//s1 := a1[:]
	//
	//// 删掉索引为1的那个3
	//s1 = append(s1[0:1], s1[2:]...)
	//fmt.Println(s1)
	//fmt.Println(a1) // ?

}
