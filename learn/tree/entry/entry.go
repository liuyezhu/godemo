package main

import (
	"fmt"
	"test_go_mod/learn/tree"
)

type MyTreeNode struct {
	*tree.Node
}

func (MyNode *MyTreeNode) postOrder() {
	if MyNode == nil || MyNode.Node == nil {
		return
	}
	left := MyTreeNode{MyNode.Left}
	right := MyTreeNode{MyNode.Right}
	left.postOrder()
	right.postOrder()
	MyNode.Print()
}

func main() {

	root := MyTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	root.postOrder()
	fmt.Println()
}
