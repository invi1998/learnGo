package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println("arr[2:6] = ", s1) //arr[2:6] =  [3 4 5 6]
	s2 := arr[:6]
	fmt.Println("arr[:6] = ", s2) // arr[:6] =  [1 2 3 4 5 6]
	s3 := arr[2:]
	fmt.Println("arr[2:] = ", s3) // arr[2:] =  [3 4 5 6 7]
	s4 := arr[:]
	fmt.Println("arr[:] = ", s4) // arr[:] =  [1 2 3 4 5 6 7]

	fmt.Println("----------------------------------------------------")

	updateSlice(s1)
	fmt.Println(s1)  // [100 4 5 6]
	fmt.Println(arr) // [1 2 100 4 5 6 7]

	updateSlice(s4)
	fmt.Println(arr) // [100 2 100 4 5 6 7]

	//Reslice
	s5 := s4[:5]
	fmt.Println(s5) // [100 2 100 4 5]
	s5 = s5[2:]
	fmt.Println(s5) // [100 4 5]

	fmt.Println("Extending slice")
	arr[0], arr[2] = 1, 3
	fmt.Println(arr) // [1 2 3 4 5 6 7]
	s1 = arr[2:6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1=[3 4 5 6], len(s1)=4, cap(s1)=5

	s2 = s1[3:5]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	//s2=[6 7], len(s2)=2, cap(s2)=2

	s3 = append(s2, 10)
	s4 = append(s3, 11)
	s5 = append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	//s3, s4, s5 =  [6 7 10] [6 7 10 11] [6 7 10 11 12]
	fmt.Println("arr = ", arr)
	//arr =  [1 2 3 4 5 6 7]
}
