package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var 数组变量名 [元素数量]Type
	// 和切片区别就是没有指定长度
	var a [10]int
	var c [10]int
	var ll [3]int
	ll[1] = 10
	fmt.Println(a, c)
	fmt.Println(reflect.TypeOf(ll))

	// 切片
	var myslice []int
	fmt.Println(reflect.TypeOf(myslice))
	b := [2]int{1, 2}
	myslice=b[:]
	fmt.Println(reflect.TypeOf(myslice))
	fmt.Println(myslice, b)
	// 定义方式二
	mysl := make([]int, 10, 10)
	fmt.Println(">>", mysl)
	fmt.Println("mysl", reflect.TypeOf(mysl))

	a1 := []int{1, 2, 3, 4, 5, 7}
	s1 := a1[:]
	fmt.Println("s1", reflect.TypeOf(s1))
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)
	//fmt.Println()

	a2 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17,13}
	fmt.Println(a2)
	s2 := a2[:]
	// 删掉索引为1的那个3
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println(a2)
	fmt.Println(s2) // ?

	s3 := make([]int, 10, 10)
	s4 := append(s3, 1)
	fmt.Println(s4)

}
