package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func LockDemo() {
	var rwMutex sync.RWMutex

	// 两个读锁
	go func() {
		fmt.Println("Read 1 started")
		rwMutex.RLock()
		time.Sleep(3 * time.Second)
		fmt.Println("Read 1 finished")
		rwMutex.RUnlock()
	}()

	go func() {
		fmt.Println("Read 2 started")
		rwMutex.RLock()
		time.Sleep(2 * time.Second)
		fmt.Println("Read 2 finished")
		rwMutex.RUnlock()
	}()
	time.Sleep(1 * time.Second)

	// 写锁
	go func() {
		rwMutex.Lock()
		fmt.Println("Write started")
		time.Sleep(3 * time.Second)
		fmt.Println("Write finished")
		rwMutex.Unlock()
	}()

	time.Sleep(5 * time.Second)
}
