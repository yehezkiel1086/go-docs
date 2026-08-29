package service

import (
	"fmt"
	"go-stacks/stack"
)

type BatchProcessor struct {
	stk stack.Stack
}

func NewBatchProcessor(stk stack.Stack) *BatchProcessor {
	return &BatchProcessor{stk: stk}
}

// ProcessBatch loads numbers into the stack and processes them in reverse order.
func (p *BatchProcessor) ProcessBatch(numbers []int) (int, error) {
	for _, n := range numbers {
		p.stk.Push(n)
	}

	sum := 0
	for !p.stk.IsEmpty() {
		val, err := p.stk.Pop()
		if err != nil {
			return 0, fmt.Errorf("failed during pop: %w", err)
		}
		sum += val
	}

	return sum, nil
}
