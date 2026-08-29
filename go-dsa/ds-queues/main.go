package main

import "fmt"

/**
Enqueue: Adds a new element to the queue.
Dequeue: Removes and returns the first (front) element from the queue.
Peek: Returns the first element in the queue.
isEmpty: Checks if the queue is empty.
Size: Finds the number of elements in the queue.
*/

type t_queue interface {
	enqueue(int)
	dequeue() int
	peek() int
	isEmpty() bool
	size() int
	String() string
}

type Queue struct {
	List []int
}

func (q *Queue) enqueue(v int) {
	q.List = append(q.List, v)
}

func (q *Queue) dequeue() int {
	tmp := q.List[0]
	q.List = q.List[1:]
	return tmp
}

func (q *Queue) peek() int {
	return q.List[0]
}

func (q *Queue) isEmpty() bool {
	return len(q.List) == 0
}

func (q *Queue) size() int {
	return len(q.List)
}

func (q *Queue) String() string {
	return fmt.Sprintf("%v", q.List)
}

func main() {
	lst := []int{1, 2, 3, 4, 5}
	var q t_queue = &Queue{
		List: lst,
	}

	// enqueue
	fmt.Println(q)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	fmt.Println(q)

	// dequeue
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	fmt.Println(q)

	// peek
	fmt.Println(q.peek())

	// isEmpty
	fmt.Println(q.isEmpty())

	// size
	fmt.Println(q.size())
}