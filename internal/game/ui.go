package game

import "github.com/hajimehoshi/ebiten/v2"

// ViewPort a player sees through
type ViewPort struct {
	Position
	width, height float64
}

// NewViewPort is initialized and returned
func NewViewPort(p Position) *ViewPort {
	w, h := ebiten.ScreenSizeInFullscreen()

	// Default scale for game is 3 px for each image px
	scale := 3

	// Scale up game for larger monitors
	if w > 1024 {
		scale = 4
	}
	if w > 1920 {
		scale = 5
	}

	vp := ViewPort{width: float64(w / scale), height: float64(h / scale)}
	vp.xPos = p.xPos
	vp.yPos = p.yPos
	return &vp
}

// Follow object
func (vp *ViewPort) Follow(o Object) {
	vp.xPos = o.xPos
	vp.yPos = o.yPos
}

// FollowAhead of object
func (vp *ViewPort) FollowAhead(o Object) {
	vp.xPos = o.xPos + (o.xSpd * (vp.width / 16))
	vp.yPos = o.yPos + (o.ySpd * (vp.height / 16))
}
