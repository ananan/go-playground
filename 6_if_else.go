package main

import "fmt"

func main() {

	// 基本使用，与for一样，判断语句不需要带小括号，但是大括号不能省
	if 5 / 2 == 0 {
		fmt.Println("5/2 = 0")
	} else {
		fmt.Println("5/2 != 0")
	}

	// if可不比配合else使用
	if 8 % 4 == 0 {
		fmt.Println("8 is even")
	}

	// 与其他语言相比，Go的if语句可以先对一个变量赋值再进行判断，这个变量的生命周期仅在if语句有效，这种语法可以使得代码更简洁
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
