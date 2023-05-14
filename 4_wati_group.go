package main

import (
	"fmt"
	"sync"
)

func hello1(num int) {
	fmt.Println("Hello Goroutine: ", num)
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer waitGroup.Done()
			hello1(j)
		}(i)
	}

	waitGroup.Wait()

}
