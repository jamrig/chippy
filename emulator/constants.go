package emulator

const (
	// MemorySize is the size of the memory in bytes.
	MemorySize = 4096
	// MemoryProgramAddress is the address at which program data starts.
	MemoryProgramAddress = 0x200
	// MemoryFontAddress is the address at which the font data starts.
	MemoryFontAddress   = 0x50
	StackInitialSize    = 32
	DisplayWidth        = 64
	DisplayHeight       = 32
	DelayTimerFrequency = 60
	CPUFrequency        = 700
)
