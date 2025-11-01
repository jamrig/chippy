package emulator

import (
	"fmt"
	"strings"
)

// TODO: handle input.

// Window represents the window which controls rendering and input.
type Window struct {
}

// NewWindow returns a new Window.
func NewWindow() *Window {
	return &Window{}
}

// Init the window.
func (w *Window) Init() {
	sb := &strings.Builder{}
	w.clear(sb)
	fmt.Print(sb.String())
}

// ShouldExit returns true if the window has requested an exit.
func (w *Window) ShouldExit() bool {
	return false
}

// Destroy frees the acquired resources.
func (w *Window) Destroy() {
}

// Render the buffer to the window.
func (w *Window) Render(buffer []byte) {
	sb := &strings.Builder{}

	w.clear(sb)

	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			if buffer[i*64+j] == 0xFF {
				w.drawWhite(sb)
			} else {
				w.drawBlack(sb)
			}
		}
		w.drawEOL(sb)
	}

	fmt.Print(sb.String())
}

func (w *Window) clear(sb *strings.Builder) {
	sb.WriteString("\033[38;5;15m\033[48;5;0m\033[H\033[2J")
}

func (w *Window) drawWhite(sb *strings.Builder) {
	sb.WriteString("\033[7m \033[27m")
}

func (w *Window) drawBlack(sb *strings.Builder) {
	sb.WriteString(" ")
}

func (w *Window) drawEOL(sb *strings.Builder) {
	sb.WriteString("\n")
}
