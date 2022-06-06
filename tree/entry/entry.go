package main

import (
	"fmt"
	"github.com/invi1998/learnGo/tree"
)

func main() {
	var root tree.Node
	fmt.Println(root)
	//	{0 <nil> <nil>}

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)
	//	[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000004078}]

	root.Print() // 3
	root.Right.Left.Setvalue(4)
	root.Right.Left.Print() // 0 值传递并不能修改到源数据

	root.Right.Left.SetvalueP(4)
	root.Right.Left.Print() // 4 引用传值就能成功修改值

	pRoot := &root
	pRoot.Print() // 3
	pRoot.SetvalueP(100)
	pRoot.Print() // 100
	root.Print()  // 100

	//var qRoot *TreeNode
	//qRoot.SetvalueP(299)
	//qRoot.Print()

	//	中序遍历
	root.TraverseMid()
	//0
	//2
	//100
	//4
	//5
	//	前序遍历
	root.TraversePre()
	//100
	//0
	//2
	//5
	//4
	//	后序遍历
	root.TraverseTil()
	//2
	//0
	//4
	//5
	//100

}