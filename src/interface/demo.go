package main

import "fmt"

//因为任何方法集都是一个空方法集的超集，所以任何类型都实现了任何空接口类型。

type Test interface {
	intertest(str string) int
	intertest2(str string) int
}

// 接口由结构体实现
type Implements struct {
}

// 实现接口 i Implements为接收者
func (i Implements) intertest(str1 string) int {
	fmt.Println(111)
	return 11
}

// 实现接口 i Implements为接收者
func (i Implements) intertest2(str1 string) int {
	fmt.Println(111)
	return 11
}

func main() {
	//demo := &Implements{}
	//demo.inertest("123")
}
