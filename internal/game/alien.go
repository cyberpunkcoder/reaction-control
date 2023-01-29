package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Alien is a squishy fusion reactor boi.
type Alien struct {
	Object
	Character
	target    *Object
	rMax      float64
	sMax      float64
	thrust    float64
	thrusting bool
}

// CreateAlien creates an alien at a position.
func CreateAlien(p Position, s Speed) *Alien {
	return &Alien{
		Object: Object{
			Position: p,
			Speed:    s,
			Image:    alienImage,
		},
		rMax:   5,
		sMax:   6,
		thrust: 0.03,
	}
}

// Update updates the the alien's state.
func (a *Alien) Update(g *Game) {
	a.NewtonsFirstLaw()

	if a.target != nil {
		a.thrusting = false

		// Direction to the target.
		dir := math.Atan2(a.target.yPos-a.yPos, a.target.xPos-a.xPos)

		if a.isGoingAwayFrom(a.target) {
			// Direction the target is moving.
			dir = math.Atan2(a.target.ySpd-a.ySpd, a.target.xSpd-a.xSpd)
		}

		// The target's rotation in degrees.
		rTarget := math.Mod(dir*(180/math.Pi)+450, 360)

		// Difference in rotation target and rotation position.
		rDiff := math.Mod((rTarget+360)-(a.rPos+360), 360)

		// If the difference in rotation is less than 45 degrees, thrust.
		if math.Abs(rDiff) < 45 {
			a.thrusting = true

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

// Draw draws the alien.
func (a *Alien) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {

	// Orient the alien's body.
	op.GeoM.Reset()
	op.GeoM.Translate(-16, -16)
	op.GeoM.Rotate(a.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(a.xPos, a.yPos)
	g.viewPort.Orient(op)

	if a.thrusting {
		// Draw the alien thrusting.
		screen.DrawImage(a.Image.SubImage(image.Rect(32, 0, 64, 32)).(*ebiten.Image), op)
	} else {
		// Draw the alien not thrusting.
		screen.DrawImage(a.Image.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), op)
	}

	// Orient the rotation of the fusion reaction inside the alien.
	op.GeoM.Reset()
	op.GeoM.Translate(-16, -16)
	op.GeoM.Rotate(math.Mod(float64(g.count)*0.2, 360) - (a.rPos / 8))
	op.GeoM.Translate(a.xPos, a.yPos)
	g.viewPort.Orient(op)
	screen.DrawImage(fusionImage, op)
}

// isGoingAwayFrom returns true if the alien is going away from the object.
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