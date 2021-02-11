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

// CreateMissile created at position with speed
func CreateMissile(p Position, s Speed) *Missile {
	// Initial missile ejection speed
	radAng := (p.rPos + 90) * (math.Pi / 180)
	s.xSpd = s.xSpd - 1*math.Cos(radAng)
	s.ySpd = s.ySpd - 1*math.Sin(radAng)

	return &Missile{
		Object: Object{
			Position: p,
			Speed:    s,
			Image:    missileImage},
		time:   0,
		delay:  25,
		burn:   50,
		thrust: 0.1,
	}
}

// Update the missile state
func (m *Missile) Update() {
	m.NewtonsFirstLaw()

	// Check if missile is active
	if m.time < m.delay+m.burn {
		if m.time >= m.delay {
			// Start thrusting
			radAng := (m.rPos + 90) * (math.Pi / 180)
			m.xSpd = m.xSpd - m.thrust*math.Cos(radAng)
			m.ySpd = m.ySpd - m.thrust*math.Sin(radAng)

			if m.time == m.delay {
				// Start thrusting sound
				queuePlayer(missile)
			}
		}
	} else if m.time == m.delay+m.burn {
		// Stop thrusting sound
		missileOff.Rewind()
		missileOff.Play()
		unQueuePlayer(missile)
	}
	m.time++
}

// Draw the missile
func (m *Missile) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	op.GeoM.Reset()
	op.GeoM.Translate(-4, -3)
	op.GeoM.Rotate(m.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(m.xPos, m.yPos)
	g.viewPort.Orient(op)

	frame := ((g.count / 2) % 2) + 1
	_, s := m.Image.Size()

	if m.time > m.delay && m.time < m.delay+m.burn {
		// Draw missile thrusting
		screen.DrawImage(m.Image.SubImage(image.Rect(frame*s, 0, s+(frame*s), s)).(*ebiten.Image), op)
		return
	}
	// Draw missile not thrusting
	screen.DrawImage(m.Image.SubImage(image.Rect(0, 0, s, s)).(*ebiten.Image), op)
}
