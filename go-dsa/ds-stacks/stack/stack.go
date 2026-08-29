package stack

type SliceStack struct {
	data []int
}

func New(data []int) *SliceStack {
	return &SliceStack{data}
}

// push
func (s *SliceStack) Push(v int) {
	s.data = append(s.data, v)
}

// peek
func (s *SliceStack) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

// pop
func (s *SliceStack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, ErrEmptyStack
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

// isEmpty
func (s *SliceStack) IsEmpty() bool {
	return s.Size() == 0
}

// Size
func (s *SliceStack) Size() int {
	return len(s.data)
}
