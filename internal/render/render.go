package render

import "github.com/hajimehoshi/ebiten/v2"

var DefaultRender *Render

type Spriter interface {
	Sprite() *ebiten.Image
}

type Render struct{}

func (r *Render) Render(screen *ebiten.Image, object Spriter) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(5, 5)
	screen.DrawImage(object.Sprite(), options)
}
