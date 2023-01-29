package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const defaultMissilesCount = 50
const defaultMaxRotationSpeed = 10
const defaultMaxSpeed = 5
const defaultDeltaV = 0.02

// Ship is a space ship.
type Ship struct {
	Object
	maxRotationSpeed          float64
	maxSpeed                  float64
	deltaV                    float64
	missileCount              int
	leftThrusters             bool
	rightThrusters            bool
	clockwiseThrusters        bool
	counterClockwiseThrusters bool
	forwardThrusters          bool
	reverseThrusters          bool
}

// CreateShip creates an initialized ship at a position with speed.
func CreateShip(p Position, s Speed) *Ship {
	return &Ship{
		Object: Object{
			Image:    shipImage,
			Position: p,
			Speed:    s,
		},
		missileCount:     defaultMissilesCount,
		maxRotationSpeed: defaultMaxRotationSpeed,
		maxSpeed:         defaultMaxSpeed,
		deltaV:           defaultDeltaV,
	}
}

// Update updates the ship's state.
func (s *Ship) Update(g *Game) {
	s.NewtonsFirstLaw()

	if s.leftThrusters {
		radAng := (s.rPos + 180) * (math.Pi / 180)
		xSpd := s.xSpd - s.deltaV*math.Cos(radAng)
		ySpd := s.ySpd - s.deltaV*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.maxSpeed {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.rightThrusters {
		radAng := (s.rPos) * (math.Pi / 180)
		xSpd := s.xSpd - s.deltaV*math.Cos(radAng)
		ySpd := s.ySpd - s.deltaV*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.maxSpeed {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.forwardThrusters {
		radAng := (s.rPos + 90) * (math.Pi / 180)
		xSpd := s.xSpd - s.deltaV*math.Cos(radAng)
		ySpd := s.ySpd - s.deltaV*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.maxSpeed {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.reverseThrusters {
		radAng := (s.rPos + 90) * (math.Pi / 180)
		xSpd := s.xSpd + s.deltaV*math.Cos(radAng)
		ySpd := s.ySpd + s.deltaV*math.Sin(radAng)

		if math.Abs(xSpd)+math.Abs(ySpd) <= s.maxSpeed {
			s.xSpd = xSpd
			s.ySpd = ySpd
		}
	}

	if s.clockwiseThrusters {
		if s.rSpd <= s.maxRotationSpeed {
			s.rSpd += s.deltaV * 2
		}
	}

	if s.counterClockwiseThrusters {
		if s.rSpd >= -s.maxRotationSpeed {
			s.rSpd -= s.deltaV * 2
		}
	}
}

// Draw draws the ship in the screen.
func (s *Ship) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	w, h := s.Image.Size()
	frame := (g.count / 2) % 2

	op.GeoM.Reset()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(s.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(s.xPos, s.yPos)
	g.viewPort.Orient(op)

	screen.DrawImage(s.Image, op)

	if s.leftThrusters {
		screen.DrawImage(rcsl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.rightThrusters {
		screen.DrawImage(rcsr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.counterClockwiseThrusters {
		screen.DrawImage(rcsfl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		screen.DrawImage(rcsbr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.clockwiseThrusters {
		screen.DrawImage(rcsfr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		screen.DrawImage(rcsbl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
	}

	if s.forwardThrusters {
		if !s.clockwiseThrusters {
			screen.DrawImage(rcsbl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}

		if !s.counterClockwiseThrusters {
			screen.DrawImage(rcsbr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}
	}

	if s.reverseThrusters {
		if !s.counterClockwiseThrusters {
			screen.DrawImage(rcsfl.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}

		if !s.clockwiseThrusters {
			screen.DrawImage(rcsfr.SubImage(image.Rect(frame*w, 0, w+(frame*w), h)).(*ebiten.Image), op)
		}
	}
}

// FireMissile fires a missile from ship.
// If ship is out of missiles, a warning sound is played.
func (s *Ship) FireMissile(g *Game) {
	// Return if ship is out of missiles
	if s.missileCount == 0 {
		if !missleEmptyPlayer.IsPlaying() {
			missleEmptyPlayer.Rewind()
			missleEmptyPlayer.Play()
		}
		return
	}
	// Missiles appear alternating from the left and right.
	offset := math.Pow(-1, float64(s.missileCount)) * 6

	pos := s.Position
	radAng := (s.rPos) * (math.Pi / 180)
	pos.xPos += offset * math.Cos(radAng)
	pos.yPos += offset * math.Sin(radAng)
	s.missileCount--

	missile := CreateMissile(pos, g.player.Speed)
	g.elements[0] = append(g.elements[0], missile)
	missileReleasePlayer.Rewind()
	missileReleasePlayer.Play()
}

// LeftThrustersOn turns the left thrusters on.
func (s *Ship) LeftThrustersOn() {
	if !s.leftThrusters {
		s.leftThrusters = true
		queuePlayer(rcsPlayer)
	}
}

// LeftThrustersOff turns the left thrusters off.
func (s *Ship) LeftThrustersOff() {
	if s.leftThrusters {
		s.leftThrusters = false
		rcsOffPlayer.Rewind()
		rcsOffPlayer.Play()
		unQueuePlayer(rcsPlayer)
	}
}

// RightThrustersOn turns the right thrusters on.
func (s *Ship) RightThrustersOn() {
	if !s.rightThrusters {
		s.rightThrusters = true
		queuePlayer(rcsPlayer)
	}
}

// RightThrustersOff turns the right thrusters off.
func (s *Ship) RightThrustersOff() {
	if s.rightThrusters {
		s.rightThrusters = false
		rcsOffPlayer.Rewind()
		rcsOffPlayer.Play()
		unQueuePlayer(rcsPlayer)
	}
}

// ClockwiseThrustersOn turns the clockwise thrusters on.
func (s *Ship) ClockwiseThrustersOn() {
	if !s.clockwiseThrusters {
		s.clockwiseThrusters = true
		queuePlayer(rcsPlayer)
	}
}

// ClockwiseThrustersOff turns the clockwise thrusters off.
func (s *Ship) ClockwiseThrustersOff() {
	if s.clockwiseThrusters {
		s.clockwiseThrusters = false
		rcsOffPlayer.Rewind()
		rcsOffPlayer.Play()
		unQueuePlayer(rcsPlayer)
	}
}

// CounterClockwiseThrustersOn turns the counter clockwise thrusters on.
func (s *Ship) CounterClockwiseThrustersOn() {
	if !s.counterClockwiseThrusters {
		s.counterClockwiseThrusters = true
		queuePlayer(rcsPlayer)
	}
}

// CounterClockwiseThrustersOff turns the counter clockwise thrusters off.
func (s *Ship) CounterClockwiseThrustersOff() {
	if s.counterClockwiseThrusters {
		s.counterClockwiseThrusters = false
		rcsOffPlayer.Rewind()
		rcsOffPlayer.Play()
		unQueuePlayer(rcsPlayer)
	}
}

// ForwardThrustersOn turns the forward thrusters on.
func (s *Ship) ForwardThrustersOn() {
	if !s.forwardThrusters {
		s.forwardThrusters = true
		queuePlayer(rcsPlayer)
	}
}

// ForwardThrustersOff turns the forward thrusters off.
func (s *Ship) ForwardThrustersOff() {
	if s.forwardThrusters {
		s.forwardThrusters = false
		rcsOffPlayer.Rewind()
		rcsOffPlayer.Play()
		unQueuePlayer(rcsPlayer)
	}
}

// ReverseThrustersOn turns the reverse thrusters on.
func (s *Ship) ReverseThrustersOn() {
	if !s.reverseThrusters {
		s.reverseThrusters = true
		queuePlayer(rcsPlayer)
	}
}

// ReverseThrustersOff turns the reverse thrusters off.
func (s *Ship) ReverseThrustersOff() {
	if s.reverseThrusters {
		s.reverseThrusters = false
		rcsOffPlayer.Rewind()
		rcsOffPlayer.Play()
		unQueuePlayer(rcsPlayer)
	}
}