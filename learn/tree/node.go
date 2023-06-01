package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {

	return &Node{Value: value}
}

func (node *Node) Print() {
	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to bil")
		return
	}
	node.Value = value
}



//func main() {
//	var root treeNode
//	root = treeNode{value: 3}
//	root.left = &treeNode{}
//	root.right = &treeNode{5, nil, nil}
//	root.right.left = new(treeNode)
//	root.left.right = createNode(2)
//
//	//nodes := []treeNode{
//	//	{value: 3},
//	//	{},
//	//	{6, nil, &root},
//	//}
//	//fmt.Println(nodes)
//	root.right.left.setValue(4)
//	root.traverse()
//	//root.right.left.print()
//	//root.print()
//	//fmt.Println()
//
//	//var pRoot *treeNode
//	//pRoot.setValue(200)
//	//pRoot = &root
//	//pRoot.setValue(100)
//	//pRoot.print()
//
//	//fmt.Println(root)
//	//fmt.Println(root.left.right)
//}
