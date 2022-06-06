package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 这种写法就是结构体方法，其实可以这样理解，就是将普通的函数形参提前，放到方法名前面，告诉编译器这是这个结构体的方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

//普通的值传递
func (node treeNode) setvalue(value int) {
	node.value = value
}

//引用传值（能够修改源）
func (node *treeNode) setvalueP(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverseMid() {
	//	中序遍历
	if node == nil {
		return
	}

	node.left.traverseMid()
	node.print()
	node.right.traverseMid()
}

func (node *treeNode) traversePre() {
	//	前序遍历
	if node == nil {
		return
	}

	node.print()
	node.left.traversePre()
	node.right.traversePre()
}

func (node *treeNode) traverseTil() {
	//	后序遍历
	if node == nil {
		return
	}

	node.left.traverseTil()
	node.right.traverseTil()
	node.print()
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
	//	go和c++这点不一样，在c++里这段代码是肯定不行的，因为它返回的是一个局部变量的地址，退出这个作用域后，这个内存会被系统回收
	//  但是go没有这个说法，go的函数里面是可以返回局部变量使用的
}

func main() {
	var root treeNode
	fmt.Println(root)
	//	{0 <nil> <nil>}

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)
	//	[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000004078}]

	root.print() // 3
	root.right.left.setvalue(4)
	root.right.left.print() // 0 值传递并不能修改到源数据

	root.right.left.setvalueP(4)
	root.right.left.print() // 4 引用传值就能成功修改值

	pRoot := &root
	pRoot.print() // 3
	pRoot.setvalueP(100)
	pRoot.print() // 100
	root.print()  // 100

	//var qRoot *treeNode
	//qRoot.setvalueP(299)
	//qRoot.print()

	//	中序遍历
	root.traverseMid()
	//0
	//2
	//100
	//4
	//5
	//	前序遍历
	root.traversePre()
	//100
	//0
	//2
	//5
	//4
	//	后序遍历
	root.traverseTil()
	//2
	//0
	//4
	//5
	//100

}
