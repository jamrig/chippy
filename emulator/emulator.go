package emulator

import (
	"os"
	"time"
)

type Emulator struct {
	Memory     *Memory
	Stack      *Stack
	DelayTimer *Timer
}

func New(fontFile, programFile string) (*Emulator, error) {
	e := &Emulator{
		Memory:     NewMemory(MemorySize),
		Stack:      NewStack(StackInitialSize),
		DelayTimer: NewTimer(DelayTimerFrequency),
	}

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

		now = time.Now().UnixNano()
		time.Sleep(1 * time.Microsecond)
	}
}

func LoadFile(file string) ([]byte, error) {
	// TODO: wrap error

	return os.ReadFile(file)
}
