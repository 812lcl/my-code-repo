package panic

import (
	"fmt"
	"os"
)

func ExitPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	os.Exit(1)
}
