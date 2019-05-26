package main

import "fmt"

func main() {

	// 基本用法，注意变量a的类型是【3】int的数组
	var a [3]int
	fmt.Println("empty: ", a)

	// 对数组中赋值可直接通过下标的方式访问
	a[1] = 100
	fmt.Println("set a: ", a)
	fmt.Println("get a: ", a[1])

	// 自带的len函数可以返回数组的长度
	fmt.Println("length: ", len(a))

	// 可以直接使用声明赋值的方式
	b := [3]int{1,2,3}
	fmt.Println(b)

	// 二维数组的用法
	var two [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two[i][j] = i + j
		}
	}
	fmt.Println("2d: ", two)
}
