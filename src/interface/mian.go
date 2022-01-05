package main

import (
	mock1 "MyGo/src/interface/mock"
	real1 "MyGo/src/interface/real"
	"fmt"
)

// 接口的实现是隐式的  只要实现里面的方法
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://wwww.baidu.com")
}

func main() {
	var r Retriever
	r = mock1.Retriever{"aa", 10}
	fmt.Printf("%T %v\n", r, r)
	r = &real1.Retriever{"mock data"}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))
}
