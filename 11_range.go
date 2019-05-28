package main

import "fmt"

func main() {

	// range对切片进行迭代，此处忽略返回对下标，使用下划线即可
	nums := []int{1,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range对map进行迭代
	dict := map[string]string{"name": "peter", "age":"23"}
	for k, v := range dict{
		fmt.Printf("%s = %s\n", k, v)
	}
}