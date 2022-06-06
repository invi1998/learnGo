package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// Print 这种写法就是结构体方法，其实可以这样理解，就是将普通的函数形参提前，放到方法名前面，告诉编译器这是这个结构体的方法
func (node Node) Print() {
	fmt.Println(node.Value)
}

// Setvalue 普通的值传递
func (node Node) Setvalue(Value int) {
	node.Value = Value
}

// SetvalueP Set valueP 引用传值（能够修改源）
func (node *Node) SetvalueP(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored")
		return
	}
	node.Value = Value
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
	//	go和c++这点不一样，在c++里这段代码是肯定不行的，因为它返回的是一个局部变量的地址，退出这个作用域后，这个内存会被系统回收
	//  但是go没有这个说法，go的函数里面是可以返回局部变量使用的
}
