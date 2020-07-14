package main

import "fmt"
type dnm func(int,int)
func main() {
	test(10,"aaaaa", "asdad")
	dnm := typefunc
	dnm(1,2)

	fmt.Println(adder(bibao(10)))
}


func test(y int, s...string) {
	println(y)
	fmt.Print(s)
}

func sl() []int{
	return nil
}

func typefunc(int, int){
	fmt.Print(111)
	prin()
}

func prin(){
	fmt.Print(2222)
}
// 闭包

func adder()(func (int) int){
	var x int
	return func(y int) int {
		return x+y
	}
}

func bibao(x int) int{
	return 10
}
