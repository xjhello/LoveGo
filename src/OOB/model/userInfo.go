package model

type userInfo struct {
	Name   string
	Age    int
	Height float32
}

type UserInfo struct {
	name   string
	age    int
	height float32
}

// 工厂模式

func NewUserInfo(name string, age int, height float32) *userInfo {
	return &userInfo{
		Name:   name,
		Age:    age,
		Height: height,
	}
}

// get 和 set  隐藏属性

func (this *userInfo) SetName(_name string) {
	this.Name = _name
}

func (this *userInfo) SetAge(_age int) {
	this.Age = _age
}

func (this *userInfo) GetAge() int {
	return this.Age
}
