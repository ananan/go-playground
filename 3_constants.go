package main

import "fmt"

const s string = "This is a Constant var"

func main() {

	//打印函数外部的静态变量
	fmt.Println(s)

	// const语句可以出现在任何var语句出现的地方
	const Pi = "3.1415926"
	fmt.Println(Pi)

	// TODO: 说明const和var语句的异同
}
