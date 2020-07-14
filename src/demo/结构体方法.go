package main

import (
	"fmt"
)

type Dog struct{
	name string
}

func NewDog(name string) *Dog{
	return &Dog{name: name}
}

func (d Dog)action() {
	fmt.Println("WAWAWAW~~~")
}

func (d *Dog) setName() {
	d.name = "BBB"
}

func main() {
	fmt.Printf("%T\n",Dog{})
	dog := NewDog("aa")
	dog.action()
	dog.setName()
	fmt.Println(dog.name)

}
