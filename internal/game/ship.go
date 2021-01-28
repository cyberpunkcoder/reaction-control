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
	missiles     int
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
		missiles: 20,
		rMax:     10,
		sMax:     5,
		thrust:   0.02,
	}
}

// Update the ship state
func (s *Ship) Update() {
	s.xPos += s.xSpd
	s.yPos += s.ySpd
	s.rPos += s.rSpd

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
}

// Draw the ship on screen in game
func (s *Ship) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	w, h := s.Image.Size()
	frame := (g.count / 2) % 2

	op.GeoM.Reset()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(s.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(s.xPos, s.yPos)
	g.viewPort.Orient(op)

	screen.DrawImage(s.Image, op)

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
	// Return if ship is out of missiles
	if s.missiles == 0 {
		if !warning.IsPlaying() {
			warning.Rewind()
			warning.Play()
		}
		return
	}
	// Alternate missiles appearing from the left and right
	offset := math.Pow(-1, float64(s.missiles)) * 6

	pos := s.Position
	radAng := (s.rPos) * (math.Pi / 180)
	pos.xPos += offset * math.Cos(radAng)
	pos.yPos += offset * math.Sin(radAng)
	s.missiles--

	missile := NewMissile(pos, g.player.Speed)
	g.elements[0] = append(g.elements[0], missile)
	release.Rewind()
	release.Play()
}

// LThrustersOn left thrusters on
func (s *Ship) LThrustersOn() {
	if !s.rThrusters {
		s.lThrusters = true
		queuePlayer(rcs)
	}
}

// LThrustersOff left thrusters off
func (s *Ship) LThrustersOff() {
	if s.lThrusters {
		s.lThrusters = false
		rcsOff.Rewind()
		rcsOff.Play()
		unQueuePlayer(rcs)
	}
}

// RThrustersOn right thrusters on
func (s *Ship) RThrustersOn() {
	if !s.rThrusters {
		s.rThrusters = true
		queuePlayer(rcs)
	}
}

// RThrustersOff right thrusters off
func (s *Ship) RThrustersOff() {
	if s.rThrusters {
		s.rThrusters = false
		rcsOff.Rewind()
		rcsOff.Play()
		unQueuePlayer(rcs)
	}
}

// CwThrustersOn clockwise thrusters on
func (s *Ship) CwThrustersOn() {
	if !s.cwThrusters {
		s.cwThrusters = true
		queuePlayer(rcs)
	}
}

// CwThrustersOff clockwise thruters off
func (s *Ship) CwThrustersOff() {
	if s.cwThrusters {
		s.cwThrusters = false
		rcsOff.Rewind()
		rcsOff.Play()
		unQueuePlayer(rcs)
	}
}

// CcwThrustersOn counter clockwise thrusters on
func (s *Ship) CcwThrustersOn() {
	if !s.ccwThrusters {
		s.ccwThrusters = true
		queuePlayer(rcs)
	}
}

// CcwThrustersOff counter clockwise thrusters off
func (s *Ship) CcwThrustersOff() {
	if s.ccwThrusters {
		s.ccwThrusters = false
		rcsOff.Rewind()
		rcsOff.Play()
		unQueuePlayer(rcs)
	}
}

// FwdThrustersOn forward thrusters on
func (s *Ship) FwdThrustersOn() {
	if !s.fwdThrusters {
		s.fwdThrusters = true
		queuePlayer(rcs)
	}
}

// FwdThrustersOff forward thrusters off
func (s *Ship) FwdThrustersOff() {
	if s.fwdThrusters {
		s.fwdThrusters = false
		rcsOff.Rewind()
		rcsOff.Play()
		unQueuePlayer(rcs)
	}
}

// RevThrustersOn reverse thrusters on
func (s *Ship) RevThrustersOn() {
	if !s.revThrusters {
		s.revThrusters = true
		queuePlayer(rcs)
	}
}

// RevThrustersOff reverse thrusters off
func (s *Ship) RevThrustersOff() {
	if s.revThrusters {
		s.revThrusters = false
		rcsOff.Rewind()
		rcsOff.Play()
		unQueuePlayer(rcs)
	}
}

func (s *Ship) isMaxSpd() bool {
	return math.Abs(s.xSpd)+math.Abs(s.ySpd) == s.sMax
}

func (s *Ship) isMaxRSpd() bool {
	return math.Abs(s.rSpd) == s.rMax
}
