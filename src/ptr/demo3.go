package main

import "fmt"

func main() {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	fmt.Printf("指针的值，也是i的值地址：%v\n", ip)
	fmt.Printf("原始指针的内存地址是：%T\n", ip)
	modify(ip)
	fmt.Println("int值被修改了，新值为:", i)
}

func modify(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	fmt.Printf("函数里接收到的指针值：%v\n", ip)
	*ip = 1
}
