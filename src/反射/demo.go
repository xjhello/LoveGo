package main

import (
	"fmt"
	"reflect"
)

//const (
//	Invalid Kind = iota  // 非法类型
//	Bool                 // 布尔型
//	Int                  // 有符号整型
//	Int8                 // 有符号8位整型
//	Int16                // 有符号16位整型
//	Int32                // 有符号32位整型
//	Int64                // 有符号64位整型
//	Uint                 // 无符号整型
//	Uint8                // 无符号8位整型
//	Uint16               // 无符号16位整型
//	Uint32               // 无符号32位整型
//	Uint64               // 无符号64位整型
//	Uintptr              // 指针
//	Float32              // 单精度浮点数
//	Float64              // 双精度浮点数
//	Complex64            // 64位复数类型
//	Complex128           // 128位复数类型
//	Array                // 数组
//	Chan                 // 通道
//	Func                 // 函数
//	Interface            // interface
//	Map                  // 映射
//	Ptr                  // 指针
//	Slice                // 切片
//	String               // 字符串
//	Struct               // 结构体
//	UnsafePointer        // 底层指针
//)
// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main() {
	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
