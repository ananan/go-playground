package main

import "fmt"

func main() {

	// 字符串可以类似Python一样用'+'相加
	fmt.Println("Hello" + "World")

	// 整型和浮点型
	fmt.Println("1+1 =", 1+1)
	fmt.Println("10.0/3.0 = ", 10.0/3.0)

	// 布尔型
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
