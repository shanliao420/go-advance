package main

import "fmt"

func main() {
	src := make(chan int)     // 无缓冲channel
	dest := make(chan int, 3) // 有缓冲的channel

	// 这里dest使用有缓冲channel主要在于消除生产者消费者执行速率的差异
	// 缓冲一下，不至于使消费者因为处理过慢影响生产者的生产速度

	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	for i := range dest {
		fmt.Println(i)
	}
}
