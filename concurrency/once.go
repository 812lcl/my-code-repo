package concurrency

import (
	"fmt"
	"sync"
)

var once sync.Once
var initializedValue int

func initialize() {
	initializedValue = 42
	fmt.Println("Initialization complete")
}

func OnceDemo() {
	// 使用 sync.Once 确保 initialize 函数只执行一次
	once.Do(initialize)

	// 无论多少次调用，initialize 函数只会执行一次
	once.Do(initialize)

	fmt.Printf("Initialized value: %d\n", initializedValue)
}
