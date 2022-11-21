package emulator

import "time"

// Timer represents a timer.
type Timer struct {
	// Value is the value of the timer.
	// This starts with being equal to the frequency and decrements.
	Value int
	// Delta is time in nanoseconds since the last update.
	Delta int64
	// UpdateDelta is the delta at which an update should happen.
	UpdateDelta int64
	// Frequency is the frequency of the timer in Hz.
	Frequency int
}

// NewTimer returns a new Timer.
func NewTimer(freq int) *Timer {
	return &Timer{
		Value:       0,
		Delta:       0,
		UpdateDelta: int64(time.Second / time.Duration(freq)),
		Frequency:   freq,
	}
}

// Tick uses the delta from the main clock cycle to update the internal state.
func (t *Timer) Tick(delta int64) {
	t.Delta += delta

	if t.Delta >= t.UpdateDelta {
		t.Delta = 0

		if t.Value == 0 {
			t.Value = t.Frequency
			return
		}

		t.Value--
	}
}

// GetValue gets the Value.
func (t *Timer) GetValue() int {
	return t.Value
}

// SetValue sets the Value.
func (t *Timer) SetValue(val int) {
	t.Value = val
}
