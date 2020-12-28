package main

import (
	"fmt"
	"unsafe"
)

type aa int32
type byte1 uint8

func swa(a,b *int){
	// t是ade1
	t := *a
	/*
	*操作符作为右值时，意义是取指针的值，作为左值时，也就是放在赋值操作符的左边时，
	表示 a 指针指向的变量。其实归纳起来，*操作符的根本意义就是操作指针指向的变量。
	当操作在右值时，就是取指向变量的值，当操作在左值时，就是将值设置给指向的变量
	 */
	*a = *b
	*b = t
}

func main() {
	var ch byte = 65
	fmt.Println(ch,unsafe.Sizeof(ch))
	var a = "asd"
	ptr := &a
	fmt.Println(unsafe.Sizeof(ptr))

	pstr := new(string)
	fmt.Println(unsafe.Sizeof(pstr))
	*pstr = "das"


	for i:=0; i<10; i++{
		fmt.Println(&i)

		fmt.Println(i)
	}
}
