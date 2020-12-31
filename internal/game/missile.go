package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Missile that the player shoots
type Missile struct {
	Object
	Physics
	Location
	time  int
	image *ebiten.Image
}

// NewMissile created at x and y coordinates
func NewMissile(location Location, physics Physics) *Missile {
	// Initial missile ejection speed
	radAng := (location.r + 90) * (math.Pi / 180)
	physics.xSpd = physics.xSpd - 1*math.Cos(radAng)
	physics.ySpd = physics.ySpd - 1*math.Sin(radAng)

	return &Missile{Location: location, Physics: physics, image: missileImage, time: 0}
}

// Update the missile state
func (missile *Missile) Update() {
	missile.x += missile.xSpd
	missile.y += missile.ySpd
	missile.r += missile.rSpd

	if missile.time < 100 {
		if missile.time > 50 {
			radAng := (missile.r + 90) * (math.Pi / 180)
			missile.xSpd = missile.xSpd - 0.04*math.Cos(radAng)
			missile.ySpd = missile.ySpd - 0.04*math.Sin(radAng)
		}
		missile.time++
	}
}

// Draw the missile
func (missile *Missile) Draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	frame := ((g.count / 2) % 2) + 1

	_, s := missile.image.Size()
	op.GeoM.Translate(-float64(s)/2, -float64(s)/2)
	op.GeoM.Rotate(float64(missile.r) * 2 * math.Pi / 360)

	x := (missile.x - g.viewPort.x) + (g.viewPort.width / 2)
	y := (missile.y - g.viewPort.y) + (g.viewPort.height / 2)

	op.GeoM.Translate(x, y)

	if missile.time > 50 && missile.time < 100 {
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

// GetPhysics of missile
func (missile *Missile) GetPhysics() Physics {
	return missile.Physics
}

// SetPhysics of missile
func (missile *Missile) SetPhysics(physics Physics) {
	missile.Physics = physics
}
