package game

import "github.com/hajimehoshi/ebiten/v2"

// Alien squishy fusion reactor boi
type Alien struct {
	Object
	Character
}

// CreateAlien at a location
func CreateAlien(p Position, s Speed) Alien {
	return Alien{Object: Object{
		Position: p,
		Speed:    s,
		Image:    alienImage,
	}}
}

// Update the alien state
func (a *Alien) Update() {

}

// Draw the alien
func (a *Alien) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {

}
