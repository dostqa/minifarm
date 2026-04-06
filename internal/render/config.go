package render

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScaleValue = 4.0
	Width      = 256
	Height     = 144
)

var Offscreen *ebiten.Image

func init() {
	Offscreen = ebiten.NewImage(Width, Height)
}
