package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Alien squishy fusion reactor boi
type Alien struct {
	Object
	Character
	target *Object
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
		rMax:   5,
		sMax:   5.5,
		thrust: 0.02,
	}
}

// Update the alien state
func (a *Alien) Update(g *Game) {
	a.NewtonsFirstLaw()

	if a.target != nil {
		// Direction to player
		dir := math.Atan2(a.target.yPos-a.yPos, a.target.xPos-a.xPos)

		if a.isGoingAwayFrom(a.target) {
			// Direction player is moving
			dir = math.Atan2(a.target.ySpd-a.ySpd, a.target.xSpd-a.xSpd)
		}

		// Target rotation degrees
		rTarget := math.Mod(dir*(180/math.Pi)+450, 360)

		// Difference in rotation target and rotation position
		rDiff := math.Mod((rTarget+360)-(a.rPos+360), 360)

		if math.Abs(rDiff) < 45 {
			// Start thrusting
			xSpd := a.xSpd + a.thrust*math.Cos(dir)
			ySpd := a.ySpd + a.thrust*math.Sin(dir)

			if math.Abs(xSpd)+math.Abs(ySpd) <= a.sMax {
				a.xSpd = xSpd
				a.ySpd = ySpd
			}
		}
		a.rSpd = rDiff / math.Pow(a.rMax, 2)
	}
}

// Draw the alien
func (a *Alien) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	var frame float64 = math.Mod(float64(g.count)*0.2, 360)

	// Orient the rotation of the fusion reaction inside the alien
	op.GeoM.Reset()
	op.GeoM.Translate(-16, -16)
	op.GeoM.Rotate(frame - (a.rPos / 4))
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

func (a *Alien) isGoingAwayFrom(o *Object) bool {
	if a.xPos > a.target.xPos {
		if a.xSpd > a.target.xSpd+a.thrust {
			return true
		}
	} else if a.xSpd < a.target.xSpd-a.thrust {
		return true
	}
	if a.yPos > a.target.yPos {
		if a.ySpd > a.target.ySpd+a.thrust {
			return true
		}
	} else if a.ySpd < a.target.ySpd-a.thrust {
		return true
	}
	return false
}
