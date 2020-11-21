package game

import "github.com/hajimehoshi/ebiten/v2"

// Objects represented in the game world
var Objects []Object

// Object in the game world
type Object interface {
	Update()
	Draw(*ebiten.Image, *ebiten.DrawImageOptions, *Game)
}
