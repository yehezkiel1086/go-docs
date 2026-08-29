package stack

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack interface {
	Push(v int)
	Peek() (int, error)
	Pop() (int, error)
	IsEmpty() bool
	Size() int
}
