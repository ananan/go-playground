// 这是学习Golang的开篇，不例外，先来个Hello World
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Hello World !")
	r := bufio.NewReader(os.Stdin)
	for {
		c, err := r.ReadString('\n')
		if err == nil {
			c = strings.Replace(c, "吗", "", -1)
			c = strings.Replace(c, "？", "！", -1)
			fmt.Println(c)
		}
	}
}
