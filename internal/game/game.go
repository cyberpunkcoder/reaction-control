package game

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	screenWidth, screenHeight, scale = 640, 480, 4
	spaceImage                       *ebiten.Image
	err                              error
)

// Game struct for ebiten
type Game struct {
	count      int
	playerShip *Ship
}

func init() {
	InitImages()
	InitSounds()
	spaceImage, _, err = ebitenutil.NewImageFromFile("../../assets/space.png")
}

func newGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.playerShip = NewShip(float64(screenWidth/2), float64(screenHeight/2))
	Objects = append(Objects, g.playerShip)
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

	for _, o := range Objects {
		o.Update()
	}

	UpdateSound()

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(spaceImage, op)

	for _, o := range Objects {
		o.Draw(screen, op, g)
	}
}

// Start the game
func (g *Game) Start() {
	ebiten.SetFullscreen(true)

	// scale up pixel art for aesthetics
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	screenWidth /= scale
	screenHeight /= scale

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
