package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Missile that the player shoots
type Missile struct {
	Object
	Physics
	Location
	time int
}

// NewMissile created at x and y coordinates
func NewMissile(x float64, y float64) *Missile {
	return &Missile{Location: Location{x: x, y: y}, time: 0}
}

// Update the missile state
func (missile *Missile) Update() {
	missile.x += missile.xSpd
	missile.y += missile.ySpd

	if missile.time < 60 {
		if missile.time > 10 {
			radAng := (missile.r + 90) * (math.Pi / 180)
			missile.xSpd = missile.xSpd - 0.02*math.Cos(radAng)
			missile.ySpd = missile.ySpd - 0.02*math.Sin(radAng)
		}
		missile.time++
	}
}

// Draw the missile
func (missile *Missile) Draw(screen *ebiten.Image, g *Game) {
	if missile.time > 10 && missile.time < 60 {
		// Draw missile thrusting
		return
	}

	// Draw missile not thrusting
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
