package main

import (
	"fmt"
)

// Go的接口是一系列方法的集合，鸭子类型
type animal interface {
	hoo() string
}

type cat struct {
	Type string
	Sound string
}

type snake struct {
	Type string
	Poisonous bool
}

// 只要实现了hoo()方法就是实现了animal接口，这个实现过程是隐式的
// 多个类型可以实现同一个接口，一个类型可以实现多个接口，是组合式的
func (s snake) hoo() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) hoo() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

// 接口类型断言


func main() {
	var a animal

	a = snake{Poisonous: true}
	fmt.Println(a.hoo())

	a = cat{Sound: "I'm a cat!"}
	fmt.Println(a.hoo())
}