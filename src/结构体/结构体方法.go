package main

import (
	"fmt"
)

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	return &Dog{name: name}
}

// 接收器
func (d Dog) action() {
	fmt.Println("WAWAWAW~~~")
}

//指针类型的接收者：类似于self或者this。可直接修改对象
func (d *Dog) setName() {
	d.name = "BBB"
}

func main() {
	fmt.Printf("%T\n", Dog{})
	dog := NewDog("aa")
	dog.action()
	dog.setName()
	fmt.Println(dog.name)

}
