package concurrency

import (
	"fmt"
	"sync"
)

func SyncMapDemo() {
	// Create a new sync.Map
	m := &sync.Map{}

	// Add some key-value pairs
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	// Retrieve a value by key
	value, ok := m.Load("key1")
	if ok {
		fmt.Println(value)
	}

	// Delete a key-value pair
	m.Delete("key2")

	// Range over all key-value pairs
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
