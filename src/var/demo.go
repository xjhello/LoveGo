package main

import "fmt"

/*
当一个变量被声明之后，系统自动赋予它该类型的零值：
int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。
所有的内存在 Go 中都是经过初始化的。
*/
func main() {
	// 变量定义格式
	// 1、标准格式 var 变量名 变量类型
	var a float32
	a = 10.1

	// 2、批量格式
	var (
		b int
		c string
	)

	// 3、简短格式 : 定义变量，同时显式初始化 只能用在函数内部
	d := 10.111
	var e = 01
	fmt.Println(a, b, c, d, e)
}
