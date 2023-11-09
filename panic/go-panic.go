package panic

import "fmt"

func PanicInGoroutine() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	go func() {
		panic("runtime error: index out of range")
	}()
}
