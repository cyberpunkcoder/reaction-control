package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const delay = 30
const burn = 50
const thrust = 0.05

// Missile that the player shoots
type Missile struct {
	Object
	time int
	//delay  int64
	//burn   int64
	//thrust int64
}

// NewMissile created at x and y coordinates
func NewMissile(location Location, speed Speed) *Missile {
	// Initial m ejection speed
	radAng := (location.r + 90) * (math.Pi / 180)
	speed.xSpd = speed.xSpd - 1*math.Cos(radAng)
	speed.ySpd = speed.ySpd - 1*math.Sin(radAng)

	return &Missile{
		Object: Object{
			Location: location,
			Speed:    speed,
			Image:    missileImage},
		time: 0}
}

// Update the m state
func (m *Missile) Update() {
	m.x += m.xSpd
	m.y += m.ySpd
	m.r += m.rSpd

	if m.time < delay+burn {
		if m.time > delay {
			radAng := (m.r + 90) * (math.Pi / 180)
			m.xSpd = m.xSpd - thrust*math.Cos(radAng)
			m.ySpd = m.ySpd - thrust*math.Sin(radAng)
		} else if m.time == delay {
			missile.Rewind()
			missile.Play()
		}
	} else if m.time == delay+burn {
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
	op.GeoM.Rotate(float64(m.r) * 2 * math.Pi / 360)

	x := (m.x - g.viewPort.x) + (g.viewPort.width / 2)
	y := (m.y - g.viewPort.y) + (g.viewPort.height / 2)

	op.GeoM.Translate(x, y)

	if m.time > delay && m.time < delay+burn {
		// Draw m thrusting
		screen.DrawImage(m.Image.SubImage(image.Rect(frame*s, 0, s+(frame*s), s)).(*ebiten.Image), op)
		return
	}
	// Draw m not thrusting
	screen.DrawImage(m.Image.SubImage(image.Rect(0, 0, s, s)).(*ebiten.Image), op)
}

// GetLocation of m
func (m *Missile) GetLocation() Location {
	return m.Location
}

// SetLocation of m
func (m *Missile) SetLocation(location Location) {
	m.Location = location
}

// GetSpeed of m
func (m *Missile) GetSpeed() Speed {
	return m.Speed
}

// SetSpeed of m
func (m *Missile) SetSpeed(speed Speed) {
	m.Speed = speed
}
