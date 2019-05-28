package main

import "fmt"

func main() {

	// 使用make创建空map
	m := make(map[string]int)

	// 对map进行set、get、len、del操作
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map: ", m)
	fmt.Println("k2: ", m["k2"])
	fmt.Println("map length: ", len(m))
	delete(m, "k2")

	// 判断某个key是否存在
	value, ok := m["k4"]
	if ok {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Key not Exists !")
	}

}
