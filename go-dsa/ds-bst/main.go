package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty = errors.New("Empty tree")
)

type TreeNode struct {
	Data        uint32
	Left, Right *TreeNode
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Printf("%v, ", root.Data)
	inorderTraversal(root.Right)
}

func search(root *TreeNode, val uint32) (*TreeNode, error) {
	if root == nil {
		return nil, ErrEmpty
	} else if root.Data == val {
		return root, nil
	} else if root.Data < val {
		return search(root.Left, val)
	} else {
		return search(root.Right, val)
	}
}

func insert(root *TreeNode, val uint32) *TreeNode {
	if root == nil {
		return &TreeNode{Data: val}
	} else {
		if val < root.Data {
			root.Left = insert(root.Left, val)
		} else if val > root.Data {
			root.Right = insert(root.Right, val)
		}
	}

	return root
}

func findLowest(root *TreeNode) uint32 {
	curr := root
	for curr.Left != nil {
		curr = curr.Left
	}

	return curr.Data
}

func deleteNode(root *TreeNode, val uint32) *TreeNode {
	if root == nil {
		return nil
	} else if val < root.Data {
		root.Left = deleteNode(root.Left, val)
	} else if val > root.Data {
		root.Right = deleteNode(root.Right, val)
	} else {
		// no left child
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			// no right child
			return root.Right
		} else {
			// with two children

			// replace current data with the lowest value in the right sub-tree
			root.Data = findLowest(root.Right)

			// delete the lowest value in the right sub-tree
			root.Right = deleteNode(root.Right, root.Data)
		}
	}
	return root
}

func main() {
	root := &TreeNode{Data: 13}
	node7 := &TreeNode{Data: 7}
	node15 := &TreeNode{Data: 15}
	node3 := &TreeNode{Data: 3}
	node8 := &TreeNode{Data: 8}
	node14 := &TreeNode{Data: 14}
	node19 := &TreeNode{Data: 19}
	node18 := &TreeNode{Data: 18}

	root.Left = node7
	root.Right = node15

	node7.Left = node3
	node7.Right = node8

	node15.Left = node14
	node15.Right = node19

	node19.Left = node18

	inorderTraversal(root)
	fmt.Println("nil")

	found, err := search(root, 13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found node: %v\n", found)
	}

	root = insert(root, 10)

	inorderTraversal(root)
	fmt.Println("nil")

	fmt.Println(findLowest(root))

	root = deleteNode(root, 15)
	inorderTraversal(root)
	fmt.Println("nil")
}
