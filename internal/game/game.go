package game

import (
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	err error
)

// Position in the game
type Position struct {
	xPos, yPos, rPos float64
}

// Speed in the game
type Speed struct {
	xSpd, ySpd, rSpd float64
}

// Element within the game
type Element interface {
	Update()
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game)
}

// Object in the game
type Object struct {
	Element
	Speed
	Position
	Mass  float64
	Image *ebiten.Image
}

// Game struct for ebiten
type Game struct {
	count    int
	player   *Ship
	viewPort *ViewPort
	elements [][]Element
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
	// Create 3 layers of objects
	// Lowest layer is for projectiles
	// Middle layer is for player and enemies
	// Highest layer is for UI
	g.elements = make([][]Element, 3)

	// Create player ship
	g.player = NewShip(Position{}, Speed{})
	g.viewPort = NewViewPort(g.player.Position)
	// Put ship on 2nd layer
	g.elements[1] = append(g.elements[1], g.player)
}

func (g *Game) control() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.player.FireMissile(g)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.player.RThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		g.player.RThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.player.LThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyE) {
		g.player.LThrustersOff()
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

// Layout the game screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.viewPort.width), int(g.viewPort.height)
}

// Update logical state of the game
func (g *Game) Update() error {
	g.count++
	g.control()

	for layer := 0; layer < len(g.elements); layer++ {
		for _, e := range g.elements[layer] {
			e.Update()
		}
	}

	g.viewPort.FollowAheadXYR(g.player.Object)
	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	w, h := space.Size()

	xMin := g.viewPort.xPos - (g.viewPort.width / 2)
	yMin := g.viewPort.yPos - (g.viewPort.height / 2)
	xMax := g.viewPort.xPos + (g.viewPort.width / 2)
	yMax := g.viewPort.yPos + (g.viewPort.height / 2)

	xMin = math.Round((xMin-float64(w))/float64(w)) * float64(w)
	xMax = math.Round((xMax+float64(w))/float64(w)) * float64(w)
	yMin = math.Round((yMin-float64(h))/float64(h)) * float64(h)
	yMax = math.Round((yMax+float64(h))/float64(h)) * float64(h)

	op := &ebiten.DrawImageOptions{}

	for x := xMin; x < xMax; x += float64(w) {
		for y := yMin; y < yMax; y += float64(h) {
			//fmt.Println(x, y)
			op.GeoM.Reset()
			op.GeoM.Translate(x, y)
			g.viewPort.Orient(op)
			screen.DrawImage(space, op)
		}
	}
	//fmt.Println()

	// Draw objects according to their layer
	for layer := 0; layer < len(g.elements); layer++ {
		for _, o := range g.elements[layer] {
			o.Draw(screen, op, g)
		}
	}

	// Testing viewport boundaries
	/*
		op.GeoM.Reset()

		w, h = shipImage.Size()
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		op.GeoM.Translate(xMin, yMin)
		g.viewPort.Orient(op)
		screen.DrawImage(shipImage, op)
		op.GeoM.Reset()

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		op.GeoM.Translate(xMax, yMax)
		g.viewPort.Orient(op)
		screen.DrawImage(shipImage, op)
	*/
}

// Start the game
func (g *Game) Start() {
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
