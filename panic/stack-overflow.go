package panic

import "fmt"

func StackOverflow() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	f()
}

func f() {
	f() // 无限递归将导致堆栈溢出
}
