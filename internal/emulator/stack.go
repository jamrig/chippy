package emulator

// Stack represents FILO stack.
type Stack struct {
	// Count is the number of items on the stack.
	Count int
	// Data is the raw stack data.
	Data []uint16
}

// NewStack creates a new Stack.
func NewStack(initialSize int) *Stack {
	return &Stack{
		Count: 0,
		Data:  make([]uint16, initialSize),
	}
}

// Pop the top item off the stack.
func (s *Stack) Pop() uint16 {
	if s.Count == 0 {
		return 0
	}

	s.Count--

	return s.Data[s.Count]
}

// Push an item onto the top of the stack.
func (s *Stack) Push(item uint16) {
	if s.Count >= len(s.Data) {
		s.Data = append(s.Data, item)
	} else {
		s.Data[s.Count] = item
	}

	s.Count++
}
