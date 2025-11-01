package emulator

// Config contains the config for the emulator.
type Config struct {
	CPU     *CPUConfig
	Memory  *MemoryConfig
	Display *DisplayConfig
}

// CPUConfig contains the config for the CPU.
type CPUConfig struct {
	// StackInitialSize is the initial size of the stack.
	StackInitialSize int
	// InstructionTimerFrequency is the frequency of the instruction timer.
	InstructionTimerFrequency int
	// DelayTimerFrequency is the frequency of the delay timer.
	DelayTimerFrequency int
	// InstructionAssignBeforeShift if true then assign Vy to Vx before shifting.
	InstructionAssignBeforeShift bool
	// InstructionUseVxForOffset if true then use Vx for offset rather than V0.
	InstructionUseVxForOffset bool
	// InstructionOverflowAddIndex if true then will set overflow flag for index addition.
	InstructionOverflowAddIndex bool
	// InstructionModifyIndexOnStoreAndLoad if true then I will be modified with Store and Load.
	InstructionModifyIndexOnStoreAndLoad bool
}

// MemoryConfig contains the config for the Memory.
type MemoryConfig struct {
	// Size is the size of the memory in bytes.
	Size uint16
	// ProgramAddress is the address at which program data starts.
	ProgramAddress uint16
	// FontAddress is the address at which the font data starts.
	FontAddress uint16
}

// DisplayConfig contains the config for the Display.
type DisplayConfig struct {
	// Width is the width of the display.
	Width int
	// Height is the height of the display.
	Height int
	// Frequency is the frequency of the display timer.
	Frequency int
}

// CHIP8Config is the base config for a CHIP-8 system (Cosmac VIP).
var CHIP8Config = &Config{
	CPU: &CPUConfig{
		StackInitialSize:                     32,
		InstructionTimerFrequency:            700,
		DelayTimerFrequency:                  60,
		InstructionAssignBeforeShift:         false,
		InstructionUseVxForOffset:            false,
		InstructionOverflowAddIndex:          false,
		InstructionModifyIndexOnStoreAndLoad: true,
	},
	Memory: &MemoryConfig{
		Size:           4096,
		ProgramAddress: 0x200,
		FontAddress:    0x50,
	},
	Display: &DisplayConfig{
		Width:     64,
		Height:    32,
		Frequency: 60,
	},
}
