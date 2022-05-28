package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 如果函数执行出错，那么把错误信息传递给error，交给使用者进行处理
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div2(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 函数可以返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

//同时函数返回可以给返回值取名,这里返回两个值，第一个返回值取名q， 第二个返回值取名r
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	//return q, r
	//	或者直接
	return
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer() // 获取函数指针
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\t", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	qd, rd := div2(22, 42)
	fmt.Println(qd, rd) // 0 22

	fmt.Println(apply(pow, 3, 4))

	//这里也可以用匿名函数
	fmt.Println(apply(
		func(i int, i2 int) int {
			return int(math.Pow(float64(i), float64(i2)))
		}, 3, 4,
	))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
