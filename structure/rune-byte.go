package structure

import "fmt"

func RuneAndByte() {
	unicodeString := "你好，世界！"

	for _, r := range unicodeString {
		fmt.Printf("%c \n", r) // 使用 rune 处理 Unicode 字符串
	}

	for i := 0; i < len(unicodeString); i++ {
		fmt.Printf("%c \n", unicodeString[i]) // 使用 byte 处理 Unicode 字符串
	}
	fmt.Println()
	asciiString := "Hello, world!"

	for _, r := range asciiString {
		fmt.Printf("%c \n", r) // 使用 rune 处理 Unicode 字符串
	}

	for i := 0; i < len(asciiString); i++ {
		fmt.Printf("%c \n", asciiString[i]) // 使用 byte 处理 ASCII 字符串
	}
	fmt.Println()
}
