package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const delay = 50
const burn = 50

// Missile that the player shoots
type Missile struct {
	Object
	Speed
	Location
	time  int
	image *ebiten.Image
}

// NewMissile created at x and y coordinates
func NewMissile(location Location, speed Speed) *Missile {
	// Initial missile ejection speed
	radAng := (location.r + 90) * (math.Pi / 180)
	speed.xSpd = speed.xSpd - 1*math.Cos(radAng)
	speed.ySpd = speed.ySpd - 1*math.Sin(radAng)

	return &Missile{Location: location, Speed: speed, image: missileImage, time: 0}
}

// Update the missile state
func (missile *Missile) Update() {
	missile.x += missile.xSpd
	missile.y += missile.ySpd
	missile.r += missile.rSpd

	if missile.time < delay+burn {
		if missile.time > delay {
			radAng := (missile.r + 90) * (math.Pi / 180)
			missile.xSpd = missile.xSpd - 0.06*math.Cos(radAng)
			missile.ySpd = missile.ySpd - 0.06*math.Sin(radAng)
		} else if missile.time == delay {
			//missileSound.Start()
		}
		missile.time++
	} else if missile.time == delay+burn {
		//missileSound.Stop()
	}
}

// Draw the missile
func (missile *Missile) Draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	frame := ((g.count / 2) % 2) + 1

	_, s := missile.image.Size()
	op.GeoM.Translate(-4, -3)
	op.GeoM.Rotate(float64(missile.r) * 2 * math.Pi / 360)

	x := (missile.x - g.viewPort.x) + (g.viewPort.width / 2)
	y := (missile.y - g.viewPort.y) + (g.viewPort.height / 2)

	op.GeoM.Translate(x, y)

	if missile.time > delay && missile.time < delay+burn {
		// Draw missile thrusting
		screen.DrawImage(missile.image.SubImage(image.Rect(frame*s, 0, s+(frame*s), s)).(*ebiten.Image), op)
		return
	}
	// Draw missile not thrusting
	screen.DrawImage(missile.image.SubImage(image.Rect(0, 0, s, s)).(*ebiten.Image), op)
}

// GetLocation of missile
func (missile *Missile) GetLocation() Location {
	return missile.Location
}

// SetLocation of missile
func (missile *Missile) SetLocation(location Location) {
	missile.Location = location
}

// GetSpeed of missile
func (missile *Missile) GetSpeed() Speed {
	return missile.Speed
}

// SetSpeed of missile
func (missile *Missile) SetSpeed(speed Speed) {
	missile.Speed = speed
}
