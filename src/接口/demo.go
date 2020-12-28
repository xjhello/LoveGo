package main

import "fmt"

type Test interface {
	intertest(str string) int
	intertest2(str string) int
}

// 接口由结构体实现
type Implements struct {
}

// 实现接口 i Implements为接收者
func (i Implements) inertest(str1 string) int {
	fmt.Println(111)
	return 11
}

// 实现接口 i Implements为接收者
func (i Implements) inertest(str1 string) int {
	fmt.Println(111)
	return 11
}

func main() {
	demo := &Implements{}
	demo.inertest("123")
}
