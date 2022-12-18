package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Missile is a missile object.
type Missile struct {
	Object
	time   int
	delay  int
	burn   int
	thrust float64
}

// CreateMissile created at a position with a speed.
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

// Update updates the missile's state.
func (m *Missile) Update(g *Game) {
	m.NewtonsFirstLaw()

	// Check if missile is active by checking if it is within it's burn time.
	if m.time < m.delay+m.burn {
		if m.time >= m.delay {
			radAng := (m.rPos + 90) * (math.Pi / 180)
			m.xSpd = m.xSpd - m.thrust*math.Cos(radAng)
			m.ySpd = m.ySpd - m.thrust*math.Sin(radAng)

			// Check if the missile is thrusting.
			if m.time == m.delay {
				queuePlayer(missile)
			}
		}
		// Check if the missile has burned out.
	} else if m.time == m.delay+m.burn {
		missileOff.Rewind()
		missileOff.Play()
		unQueuePlayer(missile)
	}
	m.time++
}

// Draw draws the missile on the screen.
func (m *Missile) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	op.GeoM.Reset()
	op.GeoM.Translate(-4, -3)
	op.GeoM.Rotate(m.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(m.xPos, m.yPos)
	g.viewPort.Orient(op)

	frame := ((g.count / 2) % 2) + 1
	_, s := m.Image.Size()

	// Check if missile is thrusting by checking if it is within it's burn time.
	if m.time > m.delay && m.time < m.delay+m.burn {
		screen.DrawImage(m.Image.SubImage(image.Rect(frame*s, 0, s+(frame*s), s)).(*ebiten.Image), op)
		return
	}
	// Draw missile but not thrusting.
	screen.DrawImage(m.Image.SubImage(image.Rect(0, 0, s, s)).(*ebiten.Image), op)
}
