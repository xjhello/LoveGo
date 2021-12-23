package main

import (
	"MyGo/src/OOB/model"
	"fmt"
)

// 对外暴露公开接口  增强安全  简化编程
func main() {
	// 小写字母开头无法被引用
	//boge := model.userInfo{
	//	Name: "1",
	//	Age: 10,
	//	Height: 10,
	//	Eduschool: "1",
	//	Hobby: []string{"coding"},
	//	MorInfo: nil,
	//}
	boge := model.NewUserInfo("haha", 10, 190.1)
	fmt.Println(boge.Name)
	// 小写字母无法实例化
	//jian := model.UserInfo{
	//	Name:   "",
	//	Age:    0,
	//	Height: 0,
	//}
	//fmt.Println(jian.Name)

	// 继承
	alipay := model.AliPay{
		PaymentArgs: model.PaymentArgs{
			AppID:       "aaa",
			MchID:       "bbb",
			Key:         "ccc",
			CallbackUrl: "ddd",
		},
		AliOpenID: "alipay123",
	}

	weixinpay := model.WeixinPay{
		PaymentArgs: model.PaymentArgs{
			AppID:       "aa12a",
			MchID:       "bbb21",
			Key:         "ccc",
			CallbackUrl: "ddd",
		},
		WeixinOpenID: "weixin123",
	}
	fmt.Println(alipay.AliOpenID)
	// 多继承 存在相同的属性 必须要知道父类
	fmt.Println(alipay.PaymentArgs.AppID)
	fmt.Println(alipay.PaymentArgs.AppID)
	fmt.Println(weixinpay.WeixinOpenID)

	alipay.Info()

	// 多态 : 1. 父类引用指向子类对象   2. 一个变量实现了接口中所有的方法 接口就可以指向这个变量
	// 接口 定义的方法不是实现. 一个类型实现了接口的所有方法  就是实现接口
	p := &payment{} // 指针类型
	p.info()

	var _pay pay
	_pay = p
	_pay.info()

	// 多重继承 所有方法都要实现
}

type A interface {
	a()
}

type B interface {
	b()
}

type C interface {
	A
	B
}

type pay interface {
	topay()
	info()
}

type payment struct {
}

func (this *payment) topay() {
	fmt.Println("topay: ")
}

func (this *payment) info() {
	fmt.Println("\n info: ")
}
