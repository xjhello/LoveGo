package main

import "fmt"

const a = 1

func swap(a, b *int)  {
	*a, *b = *b, *a
}

func printList(a [5]int)  {
	// 数组是指类型  参数传递只有值传递  会有拷贝的操作
	for i,v := range a{
		fmt.Println(i, v)
	}
}




func main() {
	// 数组
	//var testArray [3]int                        //数组会初始化为int类型的零值
	//var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	//var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化

	//var testArray [3]int
	//var numArray = [...]int{1, 2}
	//var cityArray = [...]string{"北京", "上海", "深圳"}

	//a := [...]int{1: 1, 3: 5}

	//var aList [4]int
	//var bList = [4]int{1,2}

	// 切片
	//a := []int32{1,2}
	x := 10
	y := 20
	swap(&x, &y)
	cLits := [...]int{2,3,21,3,21}
	for i:=0; i<len(cLits); i++{
		fmt.Println(cLits[i])
	}

	for i, v :=range cLits{
		fmt.Print("i:", i)
		fmt.Print("v:", v)
	}



	//切片 对比数组，没有了指定的长度限制。
	mylist := []int{1,2,}
	var aslice = []int{}
	aslice = mylist[:]
	fmt.Println("`````",aslice)
	var bslice  []string
	fmt.Println("....",len(bslice))
	//cap()函数求切片的容量
	//print(cap(aslice))

	a := make([]int, 2, 10)
	fmt.Println(a)

	//var a []string              //声明一个字符串切片
	//var b = []int{}             //声明一个整型切片并初始化
	//var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	//print(cLits)
	//print(aList, bList)
	//fmt.Print("hello")
	//a := "sadsadad"
	//fmt.Print(a)
}
