package tree

func (node *Node) TraverseMid() {
	//	中序遍历
	if node == nil {
		return
	}

	node.Left.TraverseMid()
	node.Print()
	node.Right.TraverseMid()
}

func (node *Node) TraversePre() {
	//	前序遍历
	if node == nil {
		return
	}

	node.Print()
	node.Left.TraversePre()
	node.Right.TraversePre()
}

func (node *Node) TraverseTil() {
	//	后序遍历
	if node == nil {
		return
	}

	node.Left.TraverseTil()
	node.Right.TraverseTil()
	node.Print()
}
