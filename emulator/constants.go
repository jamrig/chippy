package emulator

const (
	// MemorySize is the size of the memory in bytes.
	MemorySize = 4096
	// MemoryProgramAddress is the address at which program data starts.
	MemoryProgramAddress = 0x200
	// MemoryFontAddress is the address at which the font data starts.
	MemoryFontAddress = 0x50
	// StackInitialSize is the initial size of the stack.
	StackInitialSize = 32
	// CPUFrequency is the frequency of the CPU timer.
	CPUFrequency = 700
	// CPUUseLegacyOperations if true use the original behaviour for the CPU operations.
	CPUUseLegacyOperations = false
	// DisplayWidth is the width of the display.
	DisplayWidth = 64
	// DisplayHeight is the height of the display.
	DisplayHeight = 32
	// DisplayFrequency is the frequency of the display timer.
	DisplayFrequency = 60
	// DelayTimerFrequency is the frequency of the delay timer.
	DelayTimerFrequency = 60
)
