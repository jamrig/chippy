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
	Window     *Window
}

func New(fontFile, programFile string) (*Emulator, error) {
	font, err := LoadFile(fontFile)
	if err != nil {
		return nil, err
	}

	program, err := LoadFile(programFile)
	if err != nil {
		return nil, err
	}

	w := NewWindow()
	m := NewMemory(MemorySize)
	d := NewDisplay(DisplayWidth, DisplayHeight, DisplayFrequency, w)
	s := NewStack(StackInitialSize)
	dt := NewTimer(DelayTimerFrequency)
	e := &Emulator{
		Memory:     m,
		Stack:      s,
		DelayTimer: dt,
		CPU:        NewCPU(m, s, d, dt, CPUFrequency, MemoryProgramAddress, CPUUseLegacyOperations),
		Display:    d,
		Window:     w,
	}

	e.Memory.Write(MemoryFontAddress, font)
	e.Memory.Write(MemoryProgramAddress, program)

	return e, nil
}

func (e *Emulator) Start() {
	e.Window.Init()
	defer e.Window.Destroy()

	now := time.Now().UnixNano()
	delta := int64(0)

	for !e.Window.ShouldExit() {
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
