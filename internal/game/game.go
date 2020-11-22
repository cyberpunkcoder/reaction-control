package game

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	err error
)

// Location in the game
type Location struct {
	x, y, r float64
}

// Physics in the game
type Physics struct {
	xSpd, ySpd, rSpd, mass float64
}

// Object in the game
type Object interface {
	Update()
	Draw(*ebiten.Image, *Game)
	GetLocation() Location
	GetPhysics() Physics
}

// Game struct for ebiten
type Game struct {
	count    int
	player   *Ship
	viewPort *ViewPort
	objects  []Object
}

func init() {
	InitImages()
	InitSounds()
}

func newGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.player = NewShip(0, 0)
	g.viewPort = NewViewPort(g.player.x, g.player.y, 3)
	g.objects = append(g.objects, g.player)
}

func (g *Game) control() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.FwdThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		g.player.FwdThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.RevThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		g.player.RevThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.player.CcwThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		g.player.CcwThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.player.CwThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		g.player.CwThrustersOff()
	}
}

// Layout the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.viewPort.width), int(g.viewPort.height)
}

// Update the logical state
func (g *Game) Update() error {
	g.count++
	g.control()

	for _, o := range g.objects {
		o.Update()
	}

	g.viewPort.xPos = g.player.x + (g.player.xSpd * 80)
	g.viewPort.yPos = g.player.y + (g.player.ySpd * 80)

	UpdateSound()

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-g.viewPort.xPos, -g.viewPort.yPos)
	screen.DrawImage(space, op)

	for _, o := range g.objects {
		o.Draw(screen, g)
	}
}

// Start the game
func (g *Game) Start() {
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
