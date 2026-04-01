package render

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var DefaultRender *Render

// Spriter - любая сущность,
// которую можно вывести на экран
type Spriter interface {
	Sprite() *ebiten.Image
	X() float64
	Y() float64
}

type Render struct{}

func (r *Render) Render(screen *ebiten.Image, object Spriter) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(5, 5)
	options.GeoM.Translate(object.X(), object.Y())

	screen.DrawImage(object.Sprite(), options)
}
