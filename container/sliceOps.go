package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 1, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	//s=[2 4 6 8], len=4, cap=4

	s2 := make([]int, 16) // 创建一个len为16的slice，如果想指定slice的长度，不能直接 [16]int ，因为这样是定义一个数组
	printSlice(s2)
	//s=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16

	s3 := make([]int, 16, 32) // 创建一个len为16，cap为32的slice
	printSlice(s3)
	//s=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=32

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
	//	s=[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) // 这行语句就是将s2从下标4开始的所有元素使用功能...展开，append加到s2下标2后面（因为数组 [:3]取到下标2位置
	//	所以从效果上来看，就是删除了下标3，也就是8这个元素
	printSlice(s2)
	//	s=[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)
	//s=[4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=14, cap=15

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(front, tail) // 2 0
	printSlice(s2)
	//	s=[4 6 0 0 0 0 0 0 0 0 0 0 0], len=13, cap=15
}
