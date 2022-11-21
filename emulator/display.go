package emulator

import (
	"image"
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// White is an RGBA white.
var White = color.RGBA{R: 255, G: 255, B: 255, A: 255}

// White is an RGBA black.
var Black = color.RGBA{R: 0, G: 0, B: 0, A: 0}

// Display represents the emulator display.
type Display struct {
	Width      int
	Height     int
	Buffer     *image.RGBA
	BackBuffer *image.RGBA
	Running    bool
	Mutex      sync.Mutex
}

// NewDisplay returns a new Display.
func NewDisplay(width int, height int) *Display {
	return &Display{
		Width:      width,
		Height:     height,
		Buffer:     image.NewRGBA(image.Rect(0, 0, width, height)),
		BackBuffer: image.NewRGBA(image.Rect(0, 0, width, height)),
		Running:    false,
		Mutex:      sync.Mutex{},
	}
}

// DrawToBuffer draws the bytes to a location in the back buffer.
func (d *Display) DrawToBuffer(x int, y int, data []byte) bool {
	mx := x % d.Width
	my := y % d.Height

	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	unset := false

	for i := 0; i < len(data); i++ {
		for j := 0; j < 8; j++ {
			bit := GetBitAtPosition(data[i], 7-j)
			if bit > 0 {
				curr := d.BackBuffer.RGBAAt(mx+j, my+i)

				if curr == Black {
					d.BackBuffer.SetRGBA(mx+j, my+i, White)
				} else {
					unset = true
					d.BackBuffer.SetRGBA(mx+j, my+i, Black)
				}
			}
		}
	}

	return unset
}

// Update internal ebiten call.
func (d *Display) Update() error {
	return nil
}

// Draw internal ebiten call.
func (d *Display) Draw(screen *ebiten.Image) {
	d.Mutex.Lock()
	d.Buffer.Pix = d.BackBuffer.Pix[:]
	d.Mutex.Unlock()

	screen.WritePixels(d.BackBuffer.Pix)
}

// Layout internal ebiten call.
func (d *Display) Layout(outsideWidth, outsideHeight int) (int, int) {
	return d.Width, d.Height
}

// Start the display's main loop.
func (d *Display) Start() {
	if d.Running {
		return
	}

	d.Running = true

	ebiten.SetWindowTitle("Chippy")
	ebiten.MaximizeWindow()

	if err := ebiten.RunGame(d); err != nil {
		log.Fatal(err)
	}
}
