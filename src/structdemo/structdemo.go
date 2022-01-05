package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	city string
	age  int
}

func main() {
	// 基本实例化
	var p1 Person
	p1.name = "Tom"
	p1.age = 10
	p1.city = "beijing"
	fmt.Println(p1)

	// 匿名实例化
	var user struct {
		Name string
		age  int
	}
	fmt.Printf("%T\n", user)

	// 指针实例化
	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	var p2 = new(Person)
	var p3 = &Person{}
	fmt.Printf("%T, %T \n", p2, p3)

	// 使用键值对初始化
	p4 := Person{
		name: "Jack",
		city: "sh",
		age:  19,
	}
	fmt.Printf("%v", p4)

	p6 := &Person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6)

	// 使用值的列表初始化
	p8 := &Person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
	fmt.Println(unsafe.Sizeof(p8))

	fmt.Println(&p8)
}
