package emulator

import "time"

// DelayTimer represents a delay timer.
// The timer counts done from 60 at a frequency of 60Hz.
type DelayTimer struct {
	Value     uint8
	Delta     int64
	TickDelta int64
}

// NewDelayTimer returns a new DelayTimer.
func NewDelayTimer() *DelayTimer {
	return &DelayTimer{
		Value:     0,
		Delta:     0,
		TickDelta: int64(16 * time.Millisecond),
	}
}

// Tick uses the delta from the main clock cycle to update the internal state.
func (t *DelayTimer) Tick(delta int64) {
	t.Delta += delta

	if t.Delta >= t.TickDelta {
		t.Delta = 0

		if t.Value == 0 {
			t.Value = 60
			return
		}

		t.Value--
	}
}
