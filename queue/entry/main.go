package main

import (
	"fmt"
	"github.com/invi1998/learnGo/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}
	//    1 2 3
}
