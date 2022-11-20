package emulator

import (
	"fmt"
	"strings"
)

type Memory []byte

func NewMemory() Memory {
	m := make(Memory, MemorySize)

	m.Write(MemoryFontAddress, FontData)

	return m
}

func (m Memory) String() string {
	var b strings.Builder

	col := 0

	for i, v := range m {
		if col == 0 {
			fmt.Fprintf(&b, "\n|%08x|\t", i)
		}
		fmt.Fprintf(&b, "%02x ", v)

		col += 1

		if col == 16 {
			col = 0
		}
	}

	return b.String()
}

func (m Memory) Write(addr int, val []byte) {
	if addr < 0 || addr >= MemorySize {
		// TODO: error/logging
		return
	}

	if (len(val) + addr) > MemorySize {
		// TODO: error/logging
		return
	}

	for i, v := range val {
		m[addr+i] = v
	}
}
