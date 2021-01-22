package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Ship space ship
type Ship struct {
	Object
	rMax         float64
	sMax         float64
	lThrusters   bool
	rThrusters   bool
	cwThrusters  bool
	ccwThrusters bool
	fwdThrusters bool
	revThrusters bool
}

// NewShip is initialized and returned
func NewShip(x float64, y float64) *Ship {
	return &Ship{
		Object: Object{
			Image: shipImage,
			Location: Location{
				x: x,
				y: y,
			},
		},
		rMax: 10,
		sMax: 5,
	}
}

// Update the ship state
func (ship *Ship) Update() {
	ship.x += ship.xSpd
	ship.y += ship.ySpd
	ship.r = math.Mod(ship.r+ship.rSpd, 360)

	if ship.lThrusters {
		radAng := (ship.r + 180) * (math.Pi / 180)
		xSpd := ship.xSpd - 0.02*math.Cos(radAng)
		ySpd := ship.ySpd - 0.02*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= ship.sMax {
			ship.xSpd = xSpd
			ship.ySpd = ySpd
		}
	}

	if ship.rThrusters {
		radAng := (ship.r) * (math.Pi / 180)
		xSpd := ship.xSpd - 0.02*math.Cos(radAng)
		ySpd := ship.ySpd - 0.02*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= ship.sMax {
			ship.xSpd = xSpd
			ship.ySpd = ySpd
		}
	}

	if ship.fwdThrusters {
		radAng := (ship.r + 90) * (math.Pi / 180)
		xSpd := ship.xSpd - 0.02*math.Cos(radAng)
		ySpd := ship.ySpd - 0.02*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= ship.sMax {
			ship.xSpd = xSpd
			ship.ySpd = ySpd
		}
	}

	if ship.revThrusters {
		radAng := (ship.r + 90) * (math.Pi / 180)
		xSpd := ship.xSpd + 0.02*math.Cos(radAng)
		ySpd := ship.ySpd + 0.02*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= ship.sMax {
			ship.xSpd = xSpd
			ship.ySpd = ySpd
		}
	}

	if ship.cwThrusters {
		if ship.rSpd <= ship.rMax {
			ship.rSpd += 0.05
		}
	}

	if ship.ccwThrusters {
		if ship.rSpd >= -ship.rMax {
			ship.rSpd -= 0.05
		}
	}

	if !ship.isThrusting() {
		rcs.Pause()
	}
}

// Draw the ship on screen in game
func (ship *Ship) Draw(screen *ebiten.Image, g *Game) {

	op := &ebiten.DrawImageOptions{}

	w, h := ship.Image.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(ship.r) * 2 * math.Pi / 360)

	x := (ship.x - g.viewPort.x) + (g.viewPort.width / 2)
	y := (ship.y - g.viewPort.y) + (g.viewPort.height / 2)

	op.GeoM.Translate(x, y)
	screen.DrawImage(ship.Image, op)

	frame := (g.count / 2) % 2

	if ship.lThrusters {
		screen.DrawImage(rcsl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if ship.rThrusters {
		screen.DrawImage(rcsr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if ship.ccwThrusters {
		screen.DrawImage(rcsfl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		screen.DrawImage(rcsbr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if ship.cwThrusters {
		screen.DrawImage(rcsfr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		screen.DrawImage(rcsbl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if ship.fwdThrusters {
		if !ship.cwThrusters {
			screen.DrawImage(rcsbl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}

		if !ship.ccwThrusters {
			screen.DrawImage(rcsbr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}
	}

	if ship.revThrusters {
		if !ship.ccwThrusters {
			screen.DrawImage(rcsfl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}

		if !ship.cwThrusters {
			screen.DrawImage(rcsfr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}
	}
}

// FireMissile from ship
func (ship *Ship) FireMissile(g *Game) {
	missile := NewMissile(g.player.Location, g.player.Speed)
	g.elements[0] = append(g.elements[0], missile)
	release.Rewind()
	release.Play()
}

// GetLocation of ship
func (ship *Ship) GetLocation() Location {
	return ship.Location
}

// SetLocation of ship
func (ship *Ship) SetLocation(location Location) {
	ship.Location = location
}

// GetSpeed of ship
func (ship *Ship) GetSpeed() Speed {
	return ship.Speed
}

// SetSpeed of ship
func (ship *Ship) SetSpeed(speed Speed) {
	ship.Speed = speed
}

// LThrustersOn left thrusters on
func (ship *Ship) LThrustersOn() {
	if !ship.rThrusters && !ship.isMaxSpd() {
		ship.lThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// LThrustersOff left thrusters off
func (ship *Ship) LThrustersOff() {
	ship.lThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// RThrustersOn right thrusters on
func (ship *Ship) RThrustersOn() {
	if !ship.lThrusters && !ship.isMaxSpd() {
		ship.rThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// RThrustersOff right thrusters off
func (ship *Ship) RThrustersOff() {
	ship.rThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// CwThrustersOn clockwise thrusters on
func (ship *Ship) CwThrustersOn() {
	if !ship.cwThrusters && !ship.isMaxSpd() {
		ship.cwThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// CwThrustersOff clockwise thruters off
func (ship *Ship) CwThrustersOff() {
	ship.cwThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// CcwThrustersOn counter clockwise thrusters on
func (ship *Ship) CcwThrustersOn() {
	if !ship.ccwThrusters && !ship.isMaxSpd() {
		ship.ccwThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// CcwThrustersOff counter clockwise thrusters off
func (ship *Ship) CcwThrustersOff() {
	ship.ccwThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// FwdThrustersOn forward thrusters on
func (ship *Ship) FwdThrustersOn() {
	if !ship.fwdThrusters && !ship.isMaxSpd() {
		ship.fwdThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// FwdThrustersOff forward thrusters off
func (ship *Ship) FwdThrustersOff() {
	ship.fwdThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// RevThrustersOn reverse thrusters on
func (ship *Ship) RevThrustersOn() {
	if !ship.revThrusters && !ship.isMaxSpd() {
		ship.revThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// RevThrustersOff reverse thrusters off
func (ship *Ship) RevThrustersOff() {
	ship.revThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

func (ship *Ship) isThrusting() bool {
	return ship.lThrusters || ship.rThrusters || ship.fwdThrusters || ship.revThrusters || ship.cwThrusters || ship.ccwThrusters
}

func (ship *Ship) isMaxSpd() bool {
	return math.Abs(ship.xSpd)+math.Abs(ship.ySpd) == ship.sMax
}
