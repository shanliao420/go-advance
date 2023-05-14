package main

import (
	"sync"
	"time"
)

var (
	x     int64
	mutex sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		mutex.Lock()
		x++
		mutex.Unlock()
	}
}

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x++
	}
}

// add with lock:  10000
// add without lock:  5672

func main() {

	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("add with lock: ", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("add without lock: ", x)
}
