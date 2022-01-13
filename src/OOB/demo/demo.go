package main

type A struct {
	x string
}

func (A) y(int) bool {
	return false
}

type B struct {
	y bool
}

func (B) x(string) {}

type C struct {
	B
}

type v1 struct {
	A
	B
}

func main() {
	v := v1{}
	_ = v.A.x
}
