package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty      = errors.New("Empty Linked List")
	ErrOutOfRange = errors.New("Index out of range")
)

type Node struct {
	data uint32
	next *Node
}

func traverseLinkedList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Printf("%v -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}

func findLowest(head *Node) (uint32, error) {
	if head == nil {
		return 0, ErrEmpty
	}

	lowest := head.data
	curr := head.next
	for curr != nil {
		if curr.data < lowest {
			lowest = curr.data
		}
		curr = curr.next
	}

	return lowest, nil
}

func deleteNode(head *Node, node *Node) (*Node, error) {
	if head == nil {
		return nil, ErrEmpty
	}

	if head == node {
		return nil, nil
	}

	curr := head
	for curr.next != nil && curr.next != node {
		curr = curr.next
	}

	if curr.next != node {
		return head, nil
	}

	curr.next = curr.next.next

	return head, nil
}

func insertNodePos(head *Node, node *Node, pos int) (*Node, error) {
	if head == nil {
		return nil, ErrEmpty
	}

	if pos == 1 {
		return node, nil
	}

	curr := head
	for range pos - 2 {
		if curr == nil {
			return nil, ErrOutOfRange
		}
		curr = curr.next
	}

	node.next = curr.next
	curr.next = node

	return head, nil
}

func main() {
	node1 := &Node{data: 7}
	node2 := &Node{data: 11}
	node3 := &Node{data: 3}
	node4 := &Node{data: 2}
	node5 := &Node{data: 9}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	traverseLinkedList(node1)
	lowest, err := findLowest(node1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(lowest)
	}

	node1, err = deleteNode(node1, node4)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		traverseLinkedList(node1)
	}

	newNode := &Node{data: 97}
	node1, err = insertNodePos(node1, newNode, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		traverseLinkedList(node1)
	}
}
