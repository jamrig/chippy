package emulator

import (
	"os"
	"time"
)

// Emulator contains all of the systems for the emulator.
type Emulator struct {
	Config *Config
	Memory *Memory
	CPU    *CPU
	// Display    *Display
	// Window     *Window
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

	config := CHIP8Config

	// w := NewWindow()
	// d := NewDisplay(DisplayWidth, DisplayHeight, DisplayFrequency, w)
	e := &Emulator{
		Config: config,
		Memory: NewMemory(config.Memory.Size),

		// Display:    d,
		// Window:     w,
	}

	e.Memory.Write(config.Memory.FontAddress, font)
	e.Memory.Write(config.Memory.ProgramAddress, program)

	return e, nil
}

func (e *Emulator) Start() {
	// e.Window.Init()
	// defer e.Window.Destroy()

	now := time.Now().UnixNano()
	delta := int64(0)

	for { // !e.Window.ShouldExit() {
		delta = time.Now().UnixNano() - now

		e.CPU.Tick(delta)
		// e.Display.Tick(delta)

		now = time.Now().UnixNano()
		// TODO: use the delta to reduce the sleep
		time.Sleep(1 * time.Microsecond)
	}
}

func LoadFile(file string) ([]byte, error) {
	// TODO: wrap error

	return os.ReadFile(file)
}
