package main

import (
	"fmt"
	"time"
)

func hello(num int) {
	fmt.Println("Hello Goroutine: ", num)
}

func main() {
	for i := 0; i < 5; i++ {
		go func(j int) { // 使用go关键字 开启协程
			hello(j)
		}(i)
	}

	// 协程相较于线程更轻量级 协程 kb级别 线程 mb级别

	time.Sleep(time.Second)
}
