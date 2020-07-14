package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	//scoreMap["0"] = make(map[string]string, 10)
	fmt.Println(scoreMap["0"])
	fmt.Printf("type of a:%T\n", scoreMap)


	mymap := map[string]string{
		"name": "Tom",
		"age": "10",
	}
	fmt.Println(mymap)
	v,ok := mymap["name"]
	fmt.Println(v)
	fmt.Println(ok)


	// 报错
	var mmap map[string]string
	mmap["aaa"] = "bbb"
	fmt.Println(mmap)

}
