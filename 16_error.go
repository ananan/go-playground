package main

import (
	"errors"
	"fmt"
)

// Go的异常处理使用错误返回判断的机制
func Add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("input number is negative !")
	}
	return a + b, nil
}

func main() {
	a, b := -1, 3
	if result, err := Add(a, b); err != nil {
		fmt.Printf("Failed on add, error: %v", err)
	} else {
		fmt.Println(result)
	}

}
