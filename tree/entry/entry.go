package main

import (
	"fmt"
	"github.com/invi1998/learnGo/tree"
)

//想在Node 的基础上扩充一个后续遍历的方法
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

// 内嵌 Embedding
type myTreeNode2 struct {
	*tree.Node
}

func (myNode *myTreeNode2) postOrder2() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode2{myNode.Left}
	left.postOrder2()
	right := myTreeNode2{myNode.Right}
	right.postOrder2()
	myNode.Print()
}

// TraverseMid 重载 中序遍历
func (myNode *myTreeNode2) TraverseMid() {
	fmt.Println("this method is shadowed")
}

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

	myroot := myTreeNode{&root}
	myroot.postOrder()
	//2
	//0
	//4
	//5
	//100

	myroot2 := myTreeNode2{&root}
	myroot2.postOrder2()
	//2
	//0
	//4
	//5
	//100

	fmt.Println("--------------------------------------------------------------------------------------")
	root2 := myTreeNode2{&tree.Node{Value: 3}}
	root2.Left = &tree.Node{}
	root2.Right = &tree.Node{5, nil, nil}
	root2.Right.Left = new(tree.Node)
	root2.Left.Right = tree.CreateNode(2)
	root2.Right.Left.Setvalue(4)

	fmt.Println("In-order traversal: ")
	root2.TraverseMid()
	//In-order traversal:
	//0
	//2
	//3
	//0
	//5
	//重载之后，执行的就是
	//this method is shadowed

	//当然，这里如何跳过这层重载，去调用到源结构里的那个方法呢
	root2.Node.TraverseMid()
	//0
	//2
	//3
	//0
	//5

	fmt.Println("My own post-order traversal")
	root2.postOrder2()
	//My own post-order traversal
	//2
	//0
	//0
	//5
	//3

}
