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

// Character in the game
type Character interface {
	Up()
	Down()
	Left()
	Right()
	Cw()
	Ccw()
	Attack()
	AltAttack()
	Die()
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
	g.player = CreateShip(Position{}, Speed{})
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
	bgW, bgH := space.Size()
	// Offset each background tile one pixel to stop alias gap
	w, h := float64(bgW-1), float64(bgH-1)

	vpMaxX, vpMaxY := g.viewPort.Max()

	xMin := math.Floor((g.viewPort.xPos-vpMaxX)/w) * w
	xMax := math.Ceil((g.viewPort.xPos+vpMaxX)/w) * w
	yMin := math.Floor((g.viewPort.yPos-vpMaxY)/h) * h
	yMax := math.Ceil((g.viewPort.yPos+vpMaxY)/h) * h

	op := &ebiten.DrawImageOptions{}

	// Draw background only where needed
	for x := xMin; x < xMax; x += w {
		for y := yMin; y < yMax; y += h {
			op.GeoM.Reset()
			op.GeoM.Translate(x, y)
			g.viewPort.Orient(op)
			screen.DrawImage(space, op)
		}
	}

	// Uncomment if you want to see an alien! :)
	/*
		op.GeoM.Reset()
		op.GeoM.Translate(-16, -128)
		g.viewPort.Orient(op)
		screen.DrawImage(fusionImage, op)
		screen.DrawImage(alienImage, op)
	*/

	// Draw objects according to their layer
	for layer := 0; layer < len(g.elements); layer++ {
		for _, o := range g.elements[layer] {
			o.Draw(screen, op, g)
		}
	}
}

// Start the game
func (g *Game) Start() {
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}

// NewtonsFirstLaw states that an object will stay in motion
func (o *Object) NewtonsFirstLaw() {
	o.xPos += o.xSpd
	o.yPos += o.ySpd
	o.rPos += o.rSpd
}
