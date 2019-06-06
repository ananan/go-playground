package main

import (
	"fmt"
	"time"
)

// ”用通信来共享，而不是用共享来通信“--Go的设计哲学

// channel是Go的一个基本类型，也是实现并发的关键，主要用于协程的信息同步

func main() {
	numberChan := make(chan int)

	go producer(numberChan)

	go consumer(numberChan)

	// 上面生产者和消费者都是以协程的形式执行，此处主协程如果不等待子协程
	// 执行完毕就直接退出的话，什么都不会产生，因为协程是并发执行的
	time.Sleep(10)

}

func producer(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func consumer(ch chan int) {
	var input int
	for {
		input = <-ch
		fmt.Println(input)
	}
}