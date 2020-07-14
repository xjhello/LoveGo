package main

import "fmt"

func main() {
	//f:= func() int{
	//	return 1
	//}
	//fmt.Println(f())
	f1 := ee()
	f1("没吃饱")

	f2 := dd("匿名用户：")
	f2("有人吗？")

	f3,f4 := cc()
	f3()
	f4()
}

//func demo(tt string){
//	fmt.Println(tt)
//}

// 闭包为原函数加功能.原函数只输出一个参数。使用闭包在加一个输出
func ee() func(tt string){
	return func(tt string) {
		fmt.Println("给你加餐")
		fmt.Println(tt)
	}
}

// 闭包使用外层函数的参数

func dd(msg string) func(str string){
	return func(str string) {
		fmt.Println(msg + str)
	}
}

// 给你双倍的快乐

func cc() (func(), func()){
	func1 := func() {
		fmt.Println("1号种子选手")
	}
	func2 := func() {
		fmt.Println("2号种子选手")
	}

	return func1, func2
}