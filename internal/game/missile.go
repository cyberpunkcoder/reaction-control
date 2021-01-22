package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Missile that the player shoots
type Missile struct {
	Object
	time   int
	delay  int
	burn   int
	thrust float64
}

// NewMissile created at position with speed
func NewMissile(p Position, s Speed) *Missile {
	// Initial m ejection speed
	radAng := (p.rPos + 90) * (math.Pi / 180)
	s.xSpd = s.xSpd - 1*math.Cos(radAng)
	s.ySpd = s.ySpd - 1*math.Sin(radAng)

	return &Missile{
		Object: Object{
			Position: p,
			Speed:    s,
			Image:    missileImage},
		time:   0,
		delay:  30,
		burn:   50,
		thrust: 0.05,
	}
}

// Update the m state
func (m *Missile) Update() {
	m.xPos += m.xSpd
	m.yPos += m.ySpd
	m.rPos += m.rSpd

	// Check if missile is active
	if m.time < m.delay+m.burn {
		if m.time >= m.delay {
			// Start thrusting
			radAng := (m.rPos + 90) * (math.Pi / 180)
			m.xSpd = m.xSpd - m.thrust*math.Cos(radAng)
			m.ySpd = m.ySpd - m.thrust*math.Sin(radAng)

			if m.time == m.delay {
				// Start thrusting sound
				missile.Rewind()
				missile.Play()
			}
		}
	} else if m.time == m.delay+m.burn {
		// Stop thrusting sound
		missileOff.Rewind()
		missileOff.Play()
		missile.Pause()
	}
	m.time++
}

// Draw the m
func (m *Missile) Draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	frame := ((g.count / 2) % 2) + 1

	_, s := m.Image.Size()
	op.GeoM.Translate(-4, -3)
	op.GeoM.Rotate(float64(m.rPos) * 2 * math.Pi / 360)

	x := (m.xPos - g.viewPort.xPos) + (g.viewPort.width / 2)
	y := (m.yPos - g.viewPort.yPos) + (g.viewPort.height / 2)

	op.GeoM.Translate(x, y)

	if m.time > m.delay && m.time < m.delay+m.burn {
		// Draw m thrusting
		screen.DrawImage(m.Image.SubImage(image.Rect(frame*s, 0, s+(frame*s), s)).(*ebiten.Image), op)
		return
	}
	// Draw m not thrusting
	screen.DrawImage(m.Image.SubImage(image.Rect(0, 0, s, s)).(*ebiten.Image), op)
}
