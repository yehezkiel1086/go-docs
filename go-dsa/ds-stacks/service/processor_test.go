package service_test

import (
	"errors"
	"go-stacks/service"
	"testing"
)

// MockStack satisfies the stack.Stack interface for unit testing.
type MockStack struct {
	PushFn    func(val int)
	PopFn     func() (int, error)
	PeekFn    func() (int, error)
	IsEmptyFn func() bool
	SizeFn    func() int
}

func (m *MockStack) Push(val int) {
	if m.PushFn != nil {
		m.PushFn(val)
	}
}

func (m *MockStack) Pop() (int, error) {
	if m.PopFn != nil {
		return m.PopFn()
	}
	return 0, nil
}

func (m *MockStack) Peek() (int, error) {
	if m.PeekFn != nil {
		return m.PeekFn()
	}
	return 0, nil
}

func (m *MockStack) IsEmpty() bool {
	if m.IsEmptyFn != nil {
		return m.IsEmptyFn()
	}
	return true
}

func (m *MockStack) Size() int {
	if m.SizeFn != nil {
		return m.SizeFn()
	}
	return 0
}

func TestProcessBatch_Success(t *testing.T) {
	pushedValues := []int{}
	popCalls := 0
	simulatedStorage := []int{}

	mock := &MockStack{
		PushFn: func(val int) {
			pushedValues = append(pushedValues, val)
			simulatedStorage = append(simulatedStorage, val)
		},
		IsEmptyFn: func() bool {
			return len(simulatedStorage) == 0
		},
		PopFn: func() (int, error) {
			popCalls++
			topIdx := len(simulatedStorage) - 1
			val := simulatedStorage[topIdx]
			simulatedStorage = simulatedStorage[:topIdx]
			return val, nil
		},
	}

	processor := service.NewBatchProcessor(mock)
	result, err := processor.ProcessBatch([]int{10, 20, 30})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 60 {
		t.Errorf("expected sum 60, got %d", result)
	}

	if len(pushedValues) != 3 {
		t.Errorf("expected 3 pushes, got %d", len(pushedValues))
	}

	if popCalls != 3 {
		t.Errorf("expected 3 pops, got %d", popCalls)
	}
}

func TestProcessBatch_PopError(t *testing.T) {
	mock := &MockStack{
		PushFn: func(val int) {},
		IsEmptyFn: func() bool {
			return false // Pretend non-empty to trigger pop
		},
		PopFn: func() (int, error) {
			return 0, errors.New("underlying stack failure")
		},
	}

	processor := service.NewBatchProcessor(mock)
	_, err := processor.ProcessBatch([]int{5})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
