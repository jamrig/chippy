package emulator

// Display represents the emulator display.
type Display struct {
	Width          int
	Height         int
	Buffer         []byte
	LastTimerValue int
	Timer          *Timer
}

// NewDisplay returns a new Display.
func NewDisplay(width int, height int, freq int) *Display {
	return &Display{
		Width:          width,
		Height:         height,
		Buffer:         make([]byte, width*height),
		LastTimerValue: 0,
		Timer:          NewTimer(freq),
	}
}

// Tick calls the timer tick and if it has changed performs a display render.
func (d *Display) Tick(delta int64) {
	d.Timer.Tick(delta)

	if d.LastTimerValue != d.Timer.GetValue() {
		d.LastTimerValue = d.Timer.GetValue()

		// TODO: trigger render
	}
}

// Clear the buffer.
func (d *Display) Clear() {
	d.Buffer = make([]byte, 0)
}

// Write the bytes to a location in the buffer.
func (d *Display) Write(x int, y int, data []byte) bool {
	mx := x % d.Width
	my := y % d.Height

	unset := false

	for i := 0; i < len(data); i++ {
		if (my + i) >= d.Height {
			break
		}

		for j := 0; j < 8; j++ {
			if (mx + j) >= d.Width {
				break
			}

			bit := GetBitAtPosition(data[i], 7-j)
			if bit > 0 {
				idx := (my+i)*d.Width + mx + j

				curr := d.Buffer[idx]

				if curr == 0 {
					d.Buffer[idx] = 0xFF
				} else {
					unset = true
					d.Buffer[idx] = 0x00
				}
			}
		}
	}

	return unset
}
