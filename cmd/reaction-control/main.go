/*
author: cyberpunkprogrammer
start date: 10-30-2020
*/

package main

import (
	"os"

	"github.com/cyberpunkprogrammer/reaction-control/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	screenWidth, screenHeight, scale = 640, 480, 3
)

// Game struct for ebiten
type Game struct {
	count      int
	playerShip *world.Ship
}

func init() {
	world.InitImages()
	world.InitSounds()
}

func newGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.playerShip = world.NewShip(float64(screenWidth/2), float64(screenHeight/2))
	world.Objects = append(world.Objects, g.playerShip)
}

func (g *Game) control() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.playerShip.FwdThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		g.playerShip.FwdThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.playerShip.RevThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		g.playerShip.RevThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.playerShip.CcwThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		g.playerShip.CcwThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.playerShip.CwThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		g.playerShip.CwThrustersOff()
	}
}

// Layout the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Update the logical state
func (g *Game) Update() error {
	g.count++
	g.control()

	for _, o := range world.Objects {
		o.Update()
	}

	world.UpdateSound()

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for _, o := range world.Objects {
		o.Draw(screen, op, g.count)
	}
}

func main() {
	ebiten.SetFullscreen(true)

	// scale up pixel art for aesthetics
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	screenWidth /= scale
	screenHeight /= scale

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
