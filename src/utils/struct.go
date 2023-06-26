package utils

import "fmt"

// 定义一个结构体
type MyStruct struct {
	name string
}

// 定义结构体的方法
func (s MyStruct) MyMethod() {
	fmt.Println("Hello,", s.name)
}

// 定义一个函数
func MyFunction(s MyStruct) {
	fmt.Println("Hello from function,", s.name)
}

func Struct() {
	// 创建结构体实例
	myStruct := MyStruct{name: "John"}

	// 调用结构体方法
	myStruct.MyMethod()

	// 调用函数
	MyFunction(myStruct)
}
