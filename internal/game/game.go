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

// Speed in the game
type Speed struct {
	xSpd, ySpd, rSpd, mass float64
}

// Object in the game
type Object interface {
	Update()
	Draw(*ebiten.Image, *Game)
	GetLocation() Location
	SetLocation(Location)
	GetSpeed() Speed
	SetSpeed(Speed)
}

// Game struct for ebiten
type Game struct {
	count    int
	player   *Ship
	viewPort *ViewPort
	objects  [][]Object
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
	// Highest slayer is for UI
	g.objects = make([][]Object, 3)

	// Create player ship
	g.player = NewShip(0, 0)
	g.viewPort = NewViewPort(g.player.x, g.player.y, 4)

	// Put ship on 2nd layer
	g.objects[1] = append(g.objects[1], g.player)
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

	for layer := 0; layer < len(g.objects); layer++ {
		for _, o := range g.objects[layer] {
			o.Update()
		}
	}

	g.viewPort.FollowAhead(g.player)
	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	w, h := space.Size()
	x := (g.viewPort.x - float64(w)) - float64(int(g.viewPort.x)%w)
	y := (g.viewPort.y - float64(h)) - float64(int(g.viewPort.y)%h)

	// Draw background only where viewport is
	for i := x; i < g.viewPort.x+g.viewPort.width; i += float64(w) {
		for j := y; j < g.viewPort.y+g.viewPort.height; j += float64(h) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(i-g.viewPort.x, j-g.viewPort.y)
			screen.DrawImage(space, op)
		}
	}

	// Draw objects according to their layer
	for layer := 0; layer < len(g.objects); layer++ {
		for _, o := range g.objects[layer] {
			o.Draw(screen, g)
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
