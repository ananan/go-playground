package main

import "fmt"

func main() {

	// 最经典的for用法，类似C语言
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 替代while语句的用法
	for {
		fmt.Println("dead loop...")
		// 同样类似while， break语句用于退出循环
		break
	}

	// 带判断语句的while，实则Go已经将循环的功能都集中在for语句上了
	j := 0
	for j <= 3 {
		fmt.Println(j)
		j += 1
	}

	// 使用continue语句跳过某个循环
	for k := 0; k <= 5; k++ {
		if k % 2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
