package emulator_test

import (
	"testing"
	"time"

	"github.com/jamrig/chippy/internal/emulator"
	"github.com/stretchr/testify/assert"
)

var freq = 60
var updateDelta = int64(time.Second / time.Duration(freq))

func TestEmulator_Delay_noop(t *testing.T) {
	d := emulator.NewTimer(freq)

	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, updateDelta, d.UpdateDelta)
	assert.Equal(t, 0, d.Value)
}

func TestEmulator_Delay_partial_tick(t *testing.T) {
	d := emulator.NewTimer(freq)

	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(int64(15 * time.Millisecond))

	assert.Equal(t, int64(15*time.Millisecond), d.Delta)
	assert.Equal(t, 0, d.Value)
}

func TestEmulator_Delay_one_tick(t *testing.T) {
	d := emulator.NewTimer(freq)

	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(updateDelta)

	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 60, d.Value)
}

func TestEmulator_Delay_two_ticks(t *testing.T) {
	d := emulator.NewTimer(freq)

	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(updateDelta)
	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 60, d.Value)

	d.Tick(updateDelta)
	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 59, d.Value)
}

func TestEmulator_Delay_realistic_ticking(t *testing.T) {
	d := emulator.NewTimer(freq)

	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(int64(5 * time.Millisecond))
	assert.Equal(t, int64(5*time.Millisecond), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(int64(5 * time.Millisecond))
	assert.Equal(t, int64(10*time.Millisecond), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(int64(5 * time.Millisecond))
	assert.Equal(t, int64(15*time.Millisecond), d.Delta)
	assert.Equal(t, 0, d.Value)

	d.Tick(int64(5 * time.Millisecond))
	assert.Equal(t, int64(0), d.Delta)
	assert.Equal(t, 60, d.Value)
}
