package emulator_test

import (
	"testing"

	"github.com/jamrig/chippy/emulator"
	"github.com/stretchr/testify/assert"
)

func TestEmulator_Display_noop(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 0)

	assert.Equal(t, 10, d.Width)
	assert.Equal(t, 10, d.Height)
	assert.Equal(t, make([]byte, 100), d.Buffer)
}

func TestEmulator_Display_Clear(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 0)

	assert.Equal(t, make([]byte, 100), d.Buffer)

	d.Write(1, 5, []byte{0b11001100, 0b00110011})

	assert.NotEqual(t, make([]byte, 100), d.Buffer)

	d.Clear()

	assert.Equal(t, make([]byte, 100), d.Buffer)
}

func TestEmulator_Display_Write_in_bounds(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 0)

	assert.Equal(t, make([]byte, 100), d.Buffer)

	d.Write(1, 5, []byte{0b11001100, 0b00110011})

	assert.Equal(t, 0xFF, d.Buffer[40])
	assert.Equal(t, 0xFF, d.Buffer[41])
	assert.Equal(t, 0x00, d.Buffer[42])
	assert.Equal(t, 0x00, d.Buffer[43])
	assert.Equal(t, 0xFF, d.Buffer[44])
	assert.Equal(t, 0xFF, d.Buffer[45])
	assert.Equal(t, 0x00, d.Buffer[46])
	assert.Equal(t, 0x00, d.Buffer[47])
	assert.Equal(t, 0x00, d.Buffer[50])
	assert.Equal(t, 0x00, d.Buffer[51])
	assert.Equal(t, 0xFF, d.Buffer[52])
	assert.Equal(t, 0xFF, d.Buffer[53])
	assert.Equal(t, 0x00, d.Buffer[54])
	assert.Equal(t, 0x00, d.Buffer[55])
	assert.Equal(t, 0xFF, d.Buffer[56])
	assert.Equal(t, 0xFF, d.Buffer[57])
}

func TestEmulator_Display_Write_out_of_bounds(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 0)

	assert.Equal(t, make([]byte, 100), d.Buffer)

	d.Write(15, 19, []byte{0b11001100, 0b00110011})

	assert.Equal(t, 0xFF, d.Buffer[40])
	assert.Equal(t, 0xFF, d.Buffer[41])
	assert.Equal(t, 0x00, d.Buffer[42])
	assert.Equal(t, 0x00, d.Buffer[43])
	assert.Equal(t, 0xFF, d.Buffer[44])
	assert.Equal(t, 0x00, d.Buffer[45])
	assert.Equal(t, 0x00, d.Buffer[46])
	assert.Equal(t, 0x00, d.Buffer[47])
	assert.Equal(t, 0x00, d.Buffer[50])
	assert.Equal(t, 0x00, d.Buffer[51])
	assert.Equal(t, 0x00, d.Buffer[52])
	assert.Equal(t, 0x00, d.Buffer[53])
	assert.Equal(t, 0x00, d.Buffer[54])
	assert.Equal(t, 0x00, d.Buffer[55])
	assert.Equal(t, 0x00, d.Buffer[56])
	assert.Equal(t, 0x00, d.Buffer[57])
}
