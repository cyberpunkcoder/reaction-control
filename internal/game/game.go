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

// Position is a positon in the game.
type Position struct {
	xPos, yPos, rPos float64
}

// Speed is an x, y and rotational speed in the game.
type Speed struct {
	xSpd, ySpd, rSpd float64
}

// Element is a visual element within the game.
type Element interface {
	Update(g *Game)
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game)
}

// Character is a controllable character in the game.
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

// Object is a physical object in the game.
type Object struct {
	Element
	Speed
	Position
	Mass  float64
	Image *ebiten.Image
}

// Game is a game struct for the ebiten game engine.
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

// init creates three layers of objects in the game.
// The lowest or 0th layer is for projectiles.
// The middle or 1st layer is for player and enemies.
// The highest or 2nd layer is for UI.
func (g *Game) init() {
	g.elements = make([][]Element, 3)

	// Create the player's ship.
	g.player = CreateShip(Position{}, Speed{})
	g.viewPort = NewViewPort(g.player.Position)
	// Put the player in the middle layer.
	g.elements[1] = append(g.elements[1], g.player)

	// TODO: remove this test code.
	// Create three aliens for testing behavior.
	alien := CreateAlien(Position{0, -128, -180}, Speed{})
	alien.target = &g.player.Object
	g.elements[1] = append(g.elements[1], alien)

	alien = CreateAlien(Position{64, -128, -180}, Speed{})
	alien.target = &g.player.Object
	g.elements[1] = append(g.elements[1], alien)

	alien = CreateAlien(Position{-64, -128, -180}, Speed{})
	alien.target = &g.player.Object
	g.elements[1] = append(g.elements[1], alien)
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

// Layout is the visual layout of the game screen.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.viewPort.width), int(g.viewPort.height)
}

// Update updates the logical state of the game.indo
func (g *Game) Update() error {
	g.count++
	g.control()

	for layer := 0; layer < len(g.elements); layer++ {
		for _, e := range g.elements[layer] {
			e.Update(g)
		}
	}

	g.viewPort.FollowAheadXYR(g.player.Object)
	return nil
}

// Draw draws the game on the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	bgW, bgH := space.Size()
	// Offset each background tile one pixel to stop alias gap.
	w, h := float64(bgW-1), float64(bgH-1)

	vpMaxX, vpMaxY := g.viewPort.Max()

	xMin := math.Floor((g.viewPort.xPos-vpMaxX)/w) * w
	xMax := math.Ceil((g.viewPort.xPos+vpMaxX)/w) * w
	yMin := math.Floor((g.viewPort.yPos-vpMaxY)/h) * h
	yMax := math.Ceil((g.viewPort.yPos+vpMaxY)/h) * h

	op := &ebiten.DrawImageOptions{}

	// Draw the background only where visible by the screen.
	for x := xMin; x < xMax; x += w {
		for y := yMin; y < yMax; y += h {
			op.GeoM.Reset()
			op.GeoM.Translate(x, y)
			g.viewPort.Orient(op)
			screen.DrawImage(space, op)
		}
	}

	// Draw the objects according to their layer.
	for layer := 0; layer < len(g.elements); layer++ {
		for _, o := range g.elements[layer] {
			o.Draw(screen, op, g)
		}
	}

	// Uncomment for live debugging a value.
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", g.player.rPos))
}

// Start begins the game.
func (g *Game) Start() {
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}

// NewtonsFirstLaw applies newton's first law to an object.
// Objects in motion will stay in motion.
func (o *Object) NewtonsFirstLaw() {
	o.xPos += o.xSpd
	o.yPos += o.ySpd

	// Ensure object rotation degrees is in range 0 to 359.
	o.rPos = math.Mod(math.Abs(o.rPos+o.rSpd+360), 360)
}
