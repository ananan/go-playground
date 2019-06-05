package main

import (
	"fmt"
	"unsafe"
)

const s string = "This is a Constant var"

// 常量值还可以是len、cap、sizeof等编译器可以确定结果的函数返回值
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
)


func main() {

	//打印函数外部的静态变量
	fmt.Println(s)

	// const语句可以出现在任何var语句出现的地方
	const Pi = "3.1415926"
	fmt.Println(Pi)

	// TODO: 说明const和var语句的异同
	// fmt.Println(a) 常量值未使用不会引起编译错误，变量就会

}
