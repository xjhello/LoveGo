package main

import "fmt"

func main() {
	// &（取地址）和*（根据地址取值）。
	a := 100
	var aptr *int
	aptr = &a
	ptr := &a
	// *int  *代表指针 int代表类型  *i表示int类型变量的内存地址
	pptr := &ptr
	ppptr := &pptr
	pppptr := &ppptr
	//fmt.Println(&a)
	go fmt.Println(aptr, ptr, pptr, ppptr, pppptr)
	go fmt.Printf("%v %v %v %v \n", ptr, pptr, ppptr, pppptr)
	go fmt.Printf("%T %T %T %T \n", ptr, pptr, ppptr, pppptr)

	// new 内置函数new来为任何类型的值开辟一块内存并将此内存块的起始地址 做为此值的地址返回。
	a_new := new(int32)
	fmt.Println("a_new:", a_new)
	// 指针取值
	c := 100
	cptr := &c
	fmt.Println(*cptr)

	d := []int{1, 2, 3}
	fmt.Printf("%T \n", &d)

	//对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	//指针变量的值是指针地址。
	//对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
}
