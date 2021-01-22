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
	thrust       float64
	lThrusters   bool
	rThrusters   bool
	cwThrusters  bool
	ccwThrusters bool
	fwdThrusters bool
	revThrusters bool
}

// NewShip is initialized and returned
func NewShip(p Position, s Speed) *Ship {
	return &Ship{
		Object: Object{
			Image:    shipImage,
			Position: p,
			Speed:    s,
		},
		rMax:   10,
		sMax:   5,
		thrust: 0.02,
	}
}

// Update the ship state
func (s *Ship) Update() {
	s.xPos += s.xSpd
	s.yPos += s.ySpd
	s.rPos = math.Mod(s.rPos+s.rSpd, 360)

	if s.lThrusters {
		radAng := (s.rPos + 180) * (math.Pi / 180)
		xSpd := s.xSpd - s.thrust*math.Cos(radAng)
		ySpd := s.ySpd - s.thrust*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.sMax {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.rThrusters {
		radAng := (s.rPos) * (math.Pi / 180)
		xSpd := s.xSpd - s.thrust*math.Cos(radAng)
		ySpd := s.ySpd - s.thrust*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.sMax {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.fwdThrusters {
		radAng := (s.rPos + 90) * (math.Pi / 180)
		xSpd := s.xSpd - s.thrust*math.Cos(radAng)
		ySpd := s.ySpd - s.thrust*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.sMax {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.revThrusters {
		radAng := (s.rPos + 90) * (math.Pi / 180)
		xSpd := s.xSpd + s.thrust*math.Cos(radAng)
		ySpd := s.ySpd + s.thrust*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.sMax {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.cwThrusters {
		if s.rSpd <= s.rMax {
			s.rSpd += s.thrust * 2
		}
	}

	if s.ccwThrusters {
		if s.rSpd >= -s.rMax {
			s.rSpd -= s.thrust * 2
		}
	}

	if !s.isThrusting() {
		rcs.Pause()
	}
}

// Draw the ship on screen in game
func (s *Ship) Draw(screen *ebiten.Image, g *Game) {

	op := &ebiten.DrawImageOptions{}

	w, h := s.Image.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(s.rPos) * 2 * math.Pi / 360)

	x := (s.xPos - g.viewPort.xPos) + (g.viewPort.width / 2)
	y := (s.yPos - g.viewPort.yPos) + (g.viewPort.height / 2)

	op.GeoM.Translate(x, y)
	screen.DrawImage(s.Image, op)

	frame := (g.count / 2) % 2

	if s.lThrusters {
		screen.DrawImage(rcsl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.rThrusters {
		screen.DrawImage(rcsr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.ccwThrusters {
		screen.DrawImage(rcsfl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		screen.DrawImage(rcsbr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.cwThrusters {
		screen.DrawImage(rcsfr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		screen.DrawImage(rcsbl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.fwdThrusters {
		if !s.cwThrusters {
			screen.DrawImage(rcsbl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}

		if !s.ccwThrusters {
			screen.DrawImage(rcsbr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}
	}

	if s.revThrusters {
		if !s.ccwThrusters {
			screen.DrawImage(rcsfl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}

		if !s.cwThrusters {
			screen.DrawImage(rcsfr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}
	}
}

// FireMissile from ship
func (s *Ship) FireMissile(g *Game) {
	missile := NewMissile(g.player.Position, g.player.Speed)
	g.elements[0] = append(g.elements[0], missile)
	release.Rewind()
	release.Play()
}

// LThrustersOn left thrusters on
func (s *Ship) LThrustersOn() {
	if !s.rThrusters && !s.isMaxSpd() {
		s.lThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// LThrustersOff left thrusters off
func (s *Ship) LThrustersOff() {
	s.lThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// RThrustersOn right thrusters on
func (s *Ship) RThrustersOn() {
	if !s.lThrusters && !s.isMaxSpd() {
		s.rThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// RThrustersOff right thrusters off
func (s *Ship) RThrustersOff() {
	s.rThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// CwThrustersOn clockwise thrusters on
func (s *Ship) CwThrustersOn() {
	if !s.cwThrusters && !s.isMaxSpd() {
		s.cwThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// CwThrustersOff clockwise thruters off
func (s *Ship) CwThrustersOff() {
	s.cwThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// CcwThrustersOn counter clockwise thrusters on
func (s *Ship) CcwThrustersOn() {
	if !s.ccwThrusters && !s.isMaxSpd() {
		s.ccwThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// CcwThrustersOff counter clockwise thrusters off
func (s *Ship) CcwThrustersOff() {
	s.ccwThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// FwdThrustersOn forward thrusters on
func (s *Ship) FwdThrustersOn() {
	if !s.fwdThrusters && !s.isMaxSpd() {
		s.fwdThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// FwdThrustersOff forward thrusters off
func (s *Ship) FwdThrustersOff() {
	s.fwdThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

// RevThrustersOn reverse thrusters on
func (s *Ship) RevThrustersOn() {
	if !s.revThrusters && !s.isMaxSpd() {
		s.revThrusters = true
		rcs.Rewind()
		rcs.Play()
	}
}

// RevThrustersOff reverse thrusters off
func (s *Ship) RevThrustersOff() {
	s.revThrusters = false
	rcsOff.Rewind()
	rcsOff.Play()
}

func (s *Ship) isThrusting() bool {
	return s.lThrusters || s.rThrusters || s.fwdThrusters || s.revThrusters || s.cwThrusters || s.ccwThrusters
}

func (s *Ship) isMaxSpd() bool {
	return math.Abs(s.xSpd)+math.Abs(s.ySpd) == s.sMax
}
