package emulator

type Config struct {
	// MemorySize is the size of the memory in bytes.
	MemorySize uint16
	// MemoryProgramAddress is the address at which program data starts.
	MemoryProgramAddress uint16
	// // MemoryFontAddress is the address at which the font data starts.
	MemoryFontAddress uint16
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
	// // DisplayWidth is the width of the display.
	// DisplayWidth = 64
	// // DisplayHeight is the height of the display.
	// DisplayHeight = 32
	// // DisplayFrequency is the frequency of the display timer.
	// DisplayFrequency = 60
}

var CHIP8Config = &Config{
	MemorySize:                           4096,
	MemoryProgramAddress:                 0x200,
	MemoryFontAddress:                    0x50,
	StackInitialSize:                     32,
	InstructionTimerFrequency:            700,
	InstructionAssignBeforeShift:         false,
	InstructionUseVxForOffset:            false,
	InstructionOverflowAddIndex:          false,
	InstructionModifyIndexOnStoreAndLoad: true,
	DelayTimerFrequency:                  60,
}
