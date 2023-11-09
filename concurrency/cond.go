package concurrency

import (
	"fmt"
	"sync"
)

func CondDemo() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	data := []int{}
	maxSize := 10

	// Producer Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			mu.Lock()
			for len(data) >= maxSize {
				cond.Wait() // 如果数据已满，等待消费者信号
			}
			data = append(data, i)
			fmt.Printf("Produced: %d\n", i)
			cond.Signal() // 通知一个等待的消费者
			mu.Unlock()
		}
	}()

	// Consumer Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			mu.Lock()
			for len(data) == 0 {
				cond.Wait() // 如果数据为空，等待生产者信号
			}
			value := data[0]
			data = data[1:]
			fmt.Printf("Consumed: %d\n", value)
			cond.Signal() // 通知一个等待的生产者
			mu.Unlock()
		}
	}()
	wg.Wait()
}
