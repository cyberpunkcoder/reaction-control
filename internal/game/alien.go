package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Alien squishy fusion reactor boi
type Alien struct {
	Object
	Character
}

// CreateAlien at a location
func CreateAlien(p Position, s Speed) *Alien {
	return &Alien{Object: Object{
		Position: p,
		Speed:    s,
		Image:    alienImage,
	}}
}

// Update the alien state
func (a *Alien) Update(g *Game) {
	a.NewtonsFirstLaw()
}

// Draw the alien
func (a *Alien) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	var frame float64 = float64((g.count / 8) % 360)

	// Orient the rotation of the fusion reaction inside the alien
	op.GeoM.Reset()
	op.GeoM.Translate(-16, -16)
	op.GeoM.Rotate(frame)
	op.GeoM.Translate(a.xPos, a.yPos)
	g.viewPort.Orient(op)

	screen.DrawImage(fusionImage, op)

	// Orient the actual alien
	op.GeoM.Reset()
	op.GeoM.Translate(-16, -16)
	op.GeoM.Rotate(a.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(a.xPos, a.yPos)
	g.viewPort.Orient(op)

	screen.DrawImage(a.Image, op)
}
