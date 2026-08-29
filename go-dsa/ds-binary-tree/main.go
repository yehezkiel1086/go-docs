package main

import "fmt"

type TreeNode struct {
	Data  rune
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%c -> ", node.Data)
	preOrder(node.Left)
	preOrder(node.Right)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	fmt.Printf("%c -> ", node.Data)
	inOrder(node.Right)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	fmt.Printf("%c -> ", node.Data)
}

func main() {
	root := &TreeNode{Data: 'R'}
	nodeA := &TreeNode{Data: 'A'}
	nodeB := &TreeNode{Data: 'B'}
	nodeC := &TreeNode{Data: 'C'}
	nodeD := &TreeNode{Data: 'D'}
	nodeE := &TreeNode{Data: 'E'}
	nodeF := &TreeNode{Data: 'F'}
	nodeG := &TreeNode{Data: 'G'}

	root.Left = nodeA
	root.Right = nodeB

	nodeA.Left = nodeC
	nodeA.Right = nodeD

	nodeB.Left = nodeE
	nodeB.Right = nodeF

	nodeF.Left = nodeG

	fmt.Println("preorder:")
	preOrder(root)
	fmt.Println("nil")

	fmt.Println("inorder:")
	inOrder(root)
	fmt.Println("nil")

	fmt.Println("postorder:")
	postOrder(root)
	fmt.Println("nil")
}
