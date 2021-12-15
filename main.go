package main

import "fmt"

type Father struct {
	A string
	B string
	a string
}

func (this *Father) say(ss string) {
	fmt.Println(ss)
}

type Son struct {
	Father
	a string
}

func main() {
	//t := Son{
	//	a:"1",
	//}
	//fmt.Println(t.Father.say("11"))
}
