package main

import "fmt"

type Duck interface {
	Quack()
}
type Cat struct{}

//func (c Cat) Quack() {
//	fmt.Println("meow")
//}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

// 当我们使用指针实现接口时，只有指针类型的变量才会实现该接口；
// 当我们使用结构体实现接口时，指针类型和结构体类型都会实现该接口
func main() {
	//var c Duck = &Cat{}
	// Cat{}为值 但是接受方法中为指针  go中都是值引用 复制的未Cat的值  指向不到唯一的结构体中
	//var c Duck = Cat{}
	//c.Quack()
}
