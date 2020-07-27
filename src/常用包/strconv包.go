package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi() string转成int
	// Itoa()函数用于将int类型数据转换为对应的字符串表示
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
}
