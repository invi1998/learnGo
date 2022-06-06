package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes鼠标"

	fmt.Println(len(s)) // 9

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	//	59 65 73 E9 BC A0 E6 A0 87
	//	可以看到每个中文字符占据3个字节

	fmt.Println()

	//这里ch就是一个rune 4字节
	for i, ch := range s {
		fmt.Printf("(%d, %X)", i, ch)
	}
	//	(0, 59)(1, 65)(2, 73)(3, 9F20)(6, 6807)

	fmt.Println()

	fmt.Println("Rune count",
		utf8.RuneCountInString(s))
	//	Rune count 5

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	//	Y e s 鼠 标

	fmt.Println()

	// 使用rune转之后，会将每个字符都转成4字节长度（所以这里核上面的不同之处在于，这里不是说换一种方法去解析这段字符串的内存，而是重新为这段字符按4字节一个分配内存
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	//	(0 Y) (1 e) (2 s) (3 鼠) (4 标)

}
