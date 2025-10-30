package emulator_test

import (
	"testing"

	"github.com/jamrig/chippy/emulator"
	"github.com/stretchr/testify/assert"
)

func TestEmulator_Stack_noop(t *testing.T) {
	s := emulator.NewStack(10)

	assert.Equal(t, 0, s.Count)
	assert.Equal(t, 10, len(s.Data))
}

func TestEmulator_Stack_pop_empty(t *testing.T) {
	s := emulator.NewStack(0)

	v := s.Pop()

	assert.Equal(t, 0, s.Count)
	assert.Equal(t, 0, len(s.Data))
	assert.Equal(t, uint16(0), v)
}

func TestEmulator_Stack_push_empty(t *testing.T) {
	s := emulator.NewStack(0)

	s.Push(0xFF)

	assert.Equal(t, 1, s.Count)
	assert.Equal(t, 1, len(s.Data))
	assert.Equal(t, uint16(0xFF), s.Data[0])
}

func TestEmulator_Stack_push_pop(t *testing.T) {
	s := emulator.NewStack(1)

	s.Push(0xFF)
	v := s.Pop()

	assert.Equal(t, 0, s.Count)
	assert.Equal(t, 1, len(s.Data))
	assert.Equal(t, uint16(0xFF), v)
}

func TestEmulator_Stack_multiple(t *testing.T) {
	s := emulator.NewStack(2)

	assert.Equal(t, 0, s.Count)
	assert.Equal(t, 2, len(s.Data))

	s.Push(0x0F)
	s.Push(0xF0)
	s.Push(0xFF)
	s.Push(0xFFF)

	assert.Equal(t, 4, s.Count)
	assert.Equal(t, 4, len(s.Data))

	v0 := s.Pop()
	assert.Equal(t, 3, s.Count)

	v1 := s.Pop()
	assert.Equal(t, 2, s.Count)
	v2 := s.Pop()
	assert.Equal(t, 1, s.Count)

	s.Push(0xFFFF)
	assert.Equal(t, 2, s.Count)

	v3 := s.Pop()
	assert.Equal(t, 1, s.Count)

	v4 := s.Pop()

	assert.Equal(t, 0, s.Count)
	assert.Equal(t, 4, len(s.Data))
	assert.Equal(t, uint16(0xFFF), v0)
	assert.Equal(t, uint16(0xFF), v1)
	assert.Equal(t, uint16(0xF0), v2)
	assert.Equal(t, uint16(0xFFFF), v3)
	assert.Equal(t, uint16(0x0F), v4)
}
