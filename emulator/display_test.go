package emulator_test

import (
	"testing"

	"github.com/jamrig/chippy/emulator"
	"github.com/stretchr/testify/assert"
)

func TestEmulator_Display_noop(t *testing.T) {
	d := emulator.NewDisplay(10, 10)

	assert.Equal(t, 10, d.Width)
	assert.Equal(t, 10, d.Height)
	assert.Equal(t, make([]byte, 400), d.BackBuffer.Pix)
	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)
}

func TestEmulator_Display_Clear(t *testing.T) {
	d := emulator.NewDisplay(10, 10)

	assert.Equal(t, make([]byte, 400), d.BackBuffer.Pix)
	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)

	d.DrawToBuffer(1, 5, []byte{0b11001100, 0b00110011})

	assert.NotEqual(t, make([]byte, 400), d.BackBuffer.Pix)
	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)

	d.Clear()

	assert.Equal(t, make([]byte, 400), d.BackBuffer.Pix)
	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)
}

func TestEmulator_Display_DrawToBuffer_in_bounds(t *testing.T) {
	d := emulator.NewDisplay(10, 10)

	assert.Equal(t, make([]byte, 400), d.BackBuffer.Pix)
	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)

	d.DrawToBuffer(1, 5, []byte{0b11001100, 0b00110011})

	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(1, 5))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(2, 5))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(3, 5))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(4, 5))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(5, 5))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(6, 5))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(7, 5))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(8, 5))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(1, 6))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(2, 6))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(3, 6))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(4, 6))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(5, 6))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(6, 6))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(7, 6))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(8, 6))
}

func TestEmulator_Display_DrawToBuffer_out_of_bounds(t *testing.T) {
	d := emulator.NewDisplay(10, 10)

	assert.Equal(t, make([]byte, 400), d.BackBuffer.Pix)
	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)

	d.DrawToBuffer(15, 19, []byte{0b11001100, 0b00110011})

	assert.Equal(t, make([]byte, 400), d.Buffer.Pix)
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(5, 9))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(6, 9))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(7, 9))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(8, 9))
	assert.Equal(t, emulator.White, d.BackBuffer.RGBAAt(9, 9))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(0, 9))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(1, 9))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(2, 9))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(5, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(6, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(7, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(8, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(9, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(0, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(1, 0))
	assert.Equal(t, emulator.Black, d.BackBuffer.RGBAAt(2, 0))
}
