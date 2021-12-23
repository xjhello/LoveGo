package model

import "fmt"

type AliPay struct {
	PaymentArgs // 匿名结构体
	//  多继承相同的类
	paymentother PaymentArgsOther // 有名结构体
	AliOpenID    string
	string       // go 内置类型也可以作为属性
}

// 重写父类的Info方法
func (this *AliPay) Info() {
	fmt.Printf(">>> Info : %v", this)
}
