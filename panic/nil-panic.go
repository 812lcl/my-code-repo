package panic

import "fmt"

func NilPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	panic(nil)
}
