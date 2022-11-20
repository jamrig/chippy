package emulator

// Stack represents FILO stack.
// The stack purposely doesn't clear old data, it instead just moves the pointer.
type Stack struct {
	// Count is the number of items on the stack.
	Count int
	// Data is the raw stack data.
	Data []uint16
}

// NewStack creates a new Stack.
func NewStack(size int) *Stack {
	return &Stack{
		Count: 0,
		Data:  make([]uint16, size),
	}
}

// Pop the top item off the stack.
func (s *Stack) Pop() uint16 {
	if s.Count == 0 {
		// TODO: log
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
