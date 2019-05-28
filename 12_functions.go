package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	return a * b
}

// 多个返回值，返回值对类型需要提前指定
func addAndMulti(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, b := 3, 4

	res := add(a, b)
	fmt.Println(res)

	res = multi(a, b)
	fmt.Println(res)

	// 处理函数多返回值的时候必须返回值个数必须与函数定义的一致
	res1, res2 := addAndMulti(a, b)
	fmt.Println(res1, res2)
}

