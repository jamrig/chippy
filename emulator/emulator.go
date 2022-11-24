package emulator

import (
	"os"
	"time"
)

type Emulator struct {
	Memory     *Memory
	Stack      *Stack
	DelayTimer *Timer
	CPU        *CPU
	Display    *Display
}

func New(fontFile, programFile string) (*Emulator, error) {
	m := NewMemory(MemorySize)
	d := NewDisplay(DisplayWidth, DisplayHeight, DisplayFrequency)
	e := &Emulator{
		Memory:     m,
		Stack:      NewStack(StackInitialSize),
		DelayTimer: NewTimer(DelayTimerFrequency),
		CPU:        NewCPU(m, d, CPUFrequency, MemoryProgramAddress),
		Display:    d,
	}

	// TODO: create OpenGL
	// TODO: create input manager (using OpenGL file)
	// TODO: pass OpenGL to Display

	font, err := LoadFile(fontFile)
	if err != nil {
		return nil, err
	}

	program, err := LoadFile(programFile)
	if err != nil {
		return nil, err
	}

	e.Memory.Write(MemoryFontAddress, font)
	e.Memory.Write(MemoryProgramAddress, program)

	return e, nil
}

func (e *Emulator) Start() {
	now := time.Now().UnixNano()
	delta := int64(0)

	for {
		delta = time.Now().UnixNano() - now

		e.DelayTimer.Tick(delta)
		e.CPU.Tick(delta)
		e.Display.Tick(delta)

		now = time.Now().UnixNano()
		// TODO: use the delta to reduce the sleep
		time.Sleep(1 * time.Microsecond)
	}

}

func LoadFile(file string) ([]byte, error) {
	// TODO: wrap error

	return os.ReadFile(file)
}
