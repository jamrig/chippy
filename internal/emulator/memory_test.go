package emulator_test

import (
	"testing"

	"github.com/jamrig/chippy/internal/emulator"
	"github.com/stretchr/testify/assert"
)

func TestEmulator_Memory_noop(t *testing.T) {
	m := emulator.NewMemory(32)

	assert.Equal(t, uint16(32), m.Size)
	assert.Equal(t, make([]byte, 32), m.Data)
}

func TestEmulator_Memory_in_bounds(t *testing.T) {
	m := emulator.NewMemory(32)

	assert.Equal(t, uint16(32), m.Size)
	assert.Equal(t, make([]byte, 32), m.Data)

	m.Write(16, []byte{1, 2, 3})

	d := make([]byte, 32)
	d[16] = 1
	d[17] = 2
	d[18] = 3
	assert.Equal(t, d, m.Data)

	assert.Equal(t, byte(0), m.Read(15))
	assert.Equal(t, byte(1), m.Read(16))
	assert.Equal(t, byte(2), m.Read(17))
	assert.Equal(t, byte(3), m.Read(18))
	assert.Equal(t, byte(0), m.Read(19))
}

func TestEmulator_Memory_out_of_bounds(t *testing.T) {
	m := emulator.NewMemory(32)

	assert.Equal(t, uint16(32), m.Size)
	assert.Equal(t, make([]byte, 32), m.Data)

	m.Write(30, []byte{1, 2, 3})

	d := make([]byte, 32)
	d[30] = 1
	d[31] = 2
	assert.Equal(t, uint16(32), m.Size)
	assert.Equal(t, d, m.Data)

	assert.Equal(t, byte(0), m.Read(29))
	assert.Equal(t, byte(1), m.Read(30))
	assert.Equal(t, byte(2), m.Read(31))
	assert.Equal(t, byte(0), m.Read(32))
	assert.Equal(t, byte(0), m.Read(0))
}
