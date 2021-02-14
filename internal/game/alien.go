package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Alien squishy fusion reactor boi
type Alien struct {
	Object
	Character
	rMax   float64
	sMax   float64
	thrust float64
}

// CreateAlien at a location
func CreateAlien(p Position, s Speed) *Alien {
	return &Alien{
		Object: Object{
			Position: p,
			Speed:    s,
			Image:    alienImage,
		},
		rMax:   10,
		sMax:   5,
		thrust: 0.05,
	}
}

// Update the alien state
func (a *Alien) Update(g *Game) {
	a.NewtonsFirstLaw()

	angle := math.Atan2(g.player.yPos-a.yPos, g.player.xPos-a.xPos)*(180/math.Pi) + 90

	radAng := (angle + 90) * (math.Pi / 180)
	xSpd := a.xSpd - a.thrust*math.Cos(radAng)
	ySpd := a.ySpd - a.thrust*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) <= a.sMax {
		a.xSpd = xSpd
		a.ySpd = ySpd
	}

	a.rPos = angle
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
