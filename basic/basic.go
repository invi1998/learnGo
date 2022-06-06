package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 2, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, s := 3, 4.33, true, "def"
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, s)
}

//欧拉公式
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename)
	fmt.Println(c)
}

func enums() {
	//普通枚举类型
	const (
		red   = 0
		green = 1
		blue  = 2
	)

	//自增型枚举值
	fmt.Println(red, green, blue)

	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	fmt.Println(cpp, python, golang, javascript)

	//b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
}
