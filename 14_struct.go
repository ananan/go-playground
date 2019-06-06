package main

import "fmt"

// 结构体是不同字段的类型组合
type person struct {
	name string
	age int
	gender string
}

// 为person结构体绑定方法
func (p *person) sayHi() {
	fmt.Printf("Hello! I'm %s, %s, %v " +
		"years old.\n", p.name, p.gender, p.age)
}

func main() {

	// 创建person结构体的实例
	peter := person{name: "peter", age: 25, gender: "male"}
	fmt.Println(peter)
	fmt.Println(peter.age)
	peter.sayHi()

	// 或者直接指定值也可以创建
	savvy := person{"savvy", 24, "female"}
	fmt.Println(savvy)

	// 或使用new创建，new返回的是一个初始化置零的person对象的地址
	// 可以用下标访问实例属性，类似于其他面相对象语言的 ’类‘
	baby := new(person)
	baby.name = "Jimmy"
	fmt.Println(baby.name)
}
