package panic

import (
	"fmt"
	"runtime"
)

func GoExit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	go func() {
		defer fmt.Println("This will not be printed")
		runtime.Goexit()
	}()
}
