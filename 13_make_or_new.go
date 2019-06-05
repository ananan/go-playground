package main

import "fmt"

func main() {

	// 定义时直接提供初始化
	a := []int{0,0}
	fmt.Println(a)

	// 使用make创建，会分配内存和初始化成员结构，仅适用于slice、map、chan三种类型，返回值对象而非指针
	// make用于初始化（非零值）
	b := make([]int, 3)
	fmt.Println(b)

	// 使用new创建，适用于结构体和数组，常用语自定义结构体的初始化（置零）
	// new只接受一个参数，即类型，会计算类型大小，并为其分配零值内存，返回指针
	// new 相对来说不常用，一般的初始化用字面量或 := 就可以达到目的，除非用于自己定义的结构体
	u := new(User)
	u.age = 24
	u.name = "Peter"
	fmt.Println(u)

}

type User struct {
	name string
	age int
}