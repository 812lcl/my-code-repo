package panic

import (
	"fmt"
	"sync"
)

func RaceCondition() {
	var mu sync.Mutex
	m := make(map[string]int)

	// Wait for goroutines to finish
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		m["key"] = 42
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		value := m["key"]
		fmt.Printf("value = %v\n", value)
	}()

	wg.Wait()
}

func MuxFixRaceCondition() {
	var mu sync.Mutex
	m := make(map[string]int)

	// Wait for goroutines to finish
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		m["key"] = 42
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		value := m["key"]
		fmt.Printf("value = %v\n", value)
		mu.Unlock()
	}()

	wg.Wait()
}

func ChanToFixRaceCondition() {
	m := make(map[string]int)

	// Use channels to communicate between goroutines
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		// Write to map
		m["key"] = 42
		ch1 <- 1
	}()

	go func() {
		// Wait for write to complete
		<-ch1

		// Read from map
		value := m["key"]
		ch2 <- value
	}()

	// Wait for read to complete
	value := <-ch2
	fmt.Printf("value = %v\n", value)
}

func CondToFixRaceCondition() {
	var mu sync.Mutex
	m := make(map[string]int)

	// Create a new condition variable
	c := sync.NewCond(&mu)

	// Wait for goroutines to finish
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		m["key"] = 42
		// Signal the condition variable to wake up the other goroutine
		c.Signal()
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		// Wait for the condition variable to be signaled
		c.Wait()
		value := m["key"]
		fmt.Printf("value = %v\n", value)
		mu.Unlock()
	}()
	wg.Wait()
}
