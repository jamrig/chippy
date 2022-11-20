package main

import (
	"fmt"

	"github.com/jamrig/chippy/emulator"
)

func main() {
	m := emulator.NewMemory()
	fmt.Println(m)
}
