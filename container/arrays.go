package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}

	arr[0] = 100 // 这里把arr[0]设置为100， 看传递进来的数组是否会被改变
}

//采用指针方式来传递数组
func printArrayPointer(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}

	arr[0] = 100 // 这里把arr[0]设置为100， 看传递进来的数组是否会被改变
}

func main() {
	var arr1 [5]int                  // 定义空数组
	arr2 := [3]int{1, 2, 3}          // 使用 := 定义数组，这种定义方式（定义并初始化），需要立刻指定数组元素（初始化）
	arr3 := [...]int{2, 4, 6, 8, 10} // 使用切片来定义数组

	fmt.Println(arr1, arr2, arr3) // [0 0 0 0 0] [1 2 3] [2 4 6 8 10]

	//	二维数组
	var grid [4][5]int
	fmt.Println(grid) // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	for i := range arr3 {
		fmt.Println(i) // 0 1 2 3 4
	}

	//	同时range除了可以获取数组的下标，也可以获取数组的值
	for i, v := range arr3 {
		fmt.Println(i, v)
		/*
			0 2
			1 4
			2 6
			3 8
		*/
	}

	//	数组是值类型
	printArray(arr1)
	//printArray(arr2) // cannot use arr2 (variable of type [3]int) as type [5]int in argument to printArray
	//所以这里 [3]int 和 [5]int 是两个不同的类型
	printArray(arr3)

	fmt.Println(arr1[0]) // 0 不是 100，说明我们的数组是值类型

	printArrayPointer(&arr1)
	fmt.Println(arr1[0]) // 100

}
