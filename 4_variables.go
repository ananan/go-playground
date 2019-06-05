package main

import "fmt"

func main() {

	// var 可以声明1个或多个变量，如果声明后直接赋值那么可以不用显式知名类型，Golang会自动推导对应变量的类型
	var a string= "This is a var"
	fmt.Println(a)

	// 同时声明多个变量
	var b, c int = 1, 3
	fmt.Println(b, c)

	// 如果声明变量的时候没有赋值，系统会默认赋予对应类型的空值，比如整型为0， 布尔型则为假
	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	// 更常用的，可以使用'=:'的方式直接声明并赋值，但是值得注意的是这种写法只能在函数内部，在外部会报错
	f := "This is F var"
	fmt.Println(f)

	// 这样也是可以的
	var1, var2 := 2, 4
	fmt.Println(var1, var2)
}
