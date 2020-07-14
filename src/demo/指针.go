package main

import "fmt"

func main() {
	a := 100
	ptr := &a
	pptr := & ptr
	ppptr := & pptr
	pppptr := & ppptr
	fmt.Println(a, &a)
	fmt.Printf("%T %T %T %T \n", ptr, pptr, ppptr, pppptr)

	// 指针取值
	c := 100
	cptr := &c
	fmt.Println(*cptr)
}
