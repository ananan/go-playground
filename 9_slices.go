package main

import "fmt"

func main() {

	// 切片类似Python中的列表，如下可创建一个长度为3的string类型切片，默认初始化为零值
	s := make([]string, 3)
	fmt.Println("basic slices: ", s)

	// 可以类似数组一样使用
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	// 可以使用len获取切片长度
	fmt.Println("length: ", len(s))

	// 切片可以支持append操作，值得注意的是append是有返回值的，返回append操作后的新切片
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("After append: ", s)

	// 切片支持的切片语法类似Python，十分优雅
	newS := s[1:3]
	fmt.Println("newS: ", newS)
	fmt.Println("orig: ", s)

	// 切片下标不可以越界，用法与Python有些许差异
	newSS := s[:4]
	fmt.Println("newSS: ", newSS)

	// 或者这样
	newSSS := s[:]
	fmt.Println("newSSS: ", newSSS)

	declareSlices()
}

func declareSlices() {

	// 切片还可以这样声明
	t := []string{"a", "b", "ccccc"}
	fmt.Println("declare: ", t)
}
