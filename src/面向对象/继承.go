package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 继承Animal
}

func (d *Dog) jiao() {
	fmt.Printf("%s\n: 汪汪汪", d.name)
}

func main() {
	d1 := Dog{
		Feet: 3,
		Animal: &Animal{
			"阿黄",
		},
	}

	d1.jiao()
	d1.move()
}
