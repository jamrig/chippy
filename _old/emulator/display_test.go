package emulator_test

import (
	"testing"

	"github.com/jamrig/chippy/emulator"
	"github.com/stretchr/testify/assert"
)

func TestEmulator_Display_noop(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 1, nil)

	assert.Equal(t, 10, d.Width)
	assert.Equal(t, 10, d.Height)
	assert.Equal(t, make([]byte, 100), d.Buffer)
}

func TestEmulator_Display_Clear(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 1, nil)

	assert.Equal(t, make([]byte, 100), d.Buffer)

	d.Write(1, 5, []byte{0b11001100, 0b00110011})

	assert.NotEqual(t, make([]byte, 100), d.Buffer)

	d.Clear()

	assert.Equal(t, make([]byte, 100), d.Buffer)
}

func TestEmulator_Display_Write_in_bounds(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 1, nil)

	assert.Equal(t, make([]byte, 100), d.Buffer)

	d.Write(0, 1, []byte{0b11001100, 0b00110011})

	t.Log(d.Buffer)

	assert.Equal(t, byte(0xFF), d.Buffer[10])
	assert.Equal(t, byte(0xFF), d.Buffer[11])
	assert.Equal(t, byte(0x00), d.Buffer[12])
	assert.Equal(t, byte(0x00), d.Buffer[13])
	assert.Equal(t, byte(0xFF), d.Buffer[14])
	assert.Equal(t, byte(0xFF), d.Buffer[15])
	assert.Equal(t, byte(0x00), d.Buffer[16])
	assert.Equal(t, byte(0x00), d.Buffer[17])
	assert.Equal(t, byte(0x00), d.Buffer[20])
	assert.Equal(t, byte(0x00), d.Buffer[21])
	assert.Equal(t, byte(0xFF), d.Buffer[22])
	assert.Equal(t, byte(0xFF), d.Buffer[23])
	assert.Equal(t, byte(0x00), d.Buffer[24])
	assert.Equal(t, byte(0x00), d.Buffer[25])
	assert.Equal(t, byte(0xFF), d.Buffer[26])
	assert.Equal(t, byte(0xFF), d.Buffer[27])
}

func TestEmulator_Display_Write_out_of_bounds(t *testing.T) {
	d := emulator.NewDisplay(10, 10, 1, nil)

	assert.Equal(t, make([]byte, 100), d.Buffer)

	d.Write(10, 11, []byte{0b11001100, 0b00110011})

	assert.Equal(t, byte(0xFF), d.Buffer[10])
	assert.Equal(t, byte(0xFF), d.Buffer[11])
	assert.Equal(t, byte(0x00), d.Buffer[12])
	assert.Equal(t, byte(0x00), d.Buffer[13])
	assert.Equal(t, byte(0xFF), d.Buffer[14])
	assert.Equal(t, byte(0xFF), d.Buffer[15])
	assert.Equal(t, byte(0x00), d.Buffer[16])
	assert.Equal(t, byte(0x00), d.Buffer[17])
	assert.Equal(t, byte(0x00), d.Buffer[20])
	assert.Equal(t, byte(0x00), d.Buffer[21])
	assert.Equal(t, byte(0xFF), d.Buffer[22])
	assert.Equal(t, byte(0xFF), d.Buffer[23])
	assert.Equal(t, byte(0x00), d.Buffer[24])
	assert.Equal(t, byte(0x00), d.Buffer[25])
	assert.Equal(t, byte(0xFF), d.Buffer[26])
	assert.Equal(t, byte(0xFF), d.Buffer[27])
}
