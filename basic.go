package main

import "fmt"

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
	fmt.Printf("%T, %T, %T, %T", a, b, c, s)
}

func main() {
	fmt.Println("hello world")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
