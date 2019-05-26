package main

import (
	"fmt"
	"time"
)

func main() {

	// 基本使用, 注意case语句不需要括号
	i := 2
	fmt.Println("Write ", i, "as ")
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	}

	// case语句中可以有多个待匹配的值，使用逗号分隔开即可，default语句可以指定默认分支
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Now is Weekend !")
	default:
		fmt.Println("It's a weekday, you should go to work")
	}

	// switch不带表达式可以有if和else的功效
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 可以使用switch语句达到获取变量类型的功能
	getType("testing...")
	getType(true)
	getType(4)
}


// 获取变量类型，直接打印出来
func getType(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("It's a bool")
	case int:
		fmt.Println("It's a int")
	case string:
		fmt.Println("It's string")
	default:
		fmt.Println("Don't Know")

	}
}
