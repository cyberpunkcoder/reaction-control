package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// ViewPort a player sees through
type ViewPort struct {
	Position
	Speed
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

// Orient drawing options according to viewport
func (vp *ViewPort) Orient(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(-vp.xPos, -vp.yPos)
	op.GeoM.Rotate(-vp.rPos * 2 * math.Pi / 360)
	op.GeoM.Translate(vp.width/2, vp.height/2)
}

// LockXY position of viewport to object
func (vp *ViewPort) LockXY(o Object) {
	vp.xPos = o.xPos
	vp.yPos = o.yPos
}

// LockXYR position of viewport to object
func (vp *ViewPort) LockXYR(o Object) {
	vp.xPos = o.xPos
	vp.yPos = o.yPos
	vp.rPos = o.rPos
}

// FollowAheadXY position of object
func (vp *ViewPort) FollowAheadXY(o Object) {
	vp.xPos = o.xPos + (o.xSpd * (vp.width / 16))
	vp.yPos = o.yPos + (o.ySpd * (vp.height / 16))
}

// FollowAheadXYR position of object
func (vp *ViewPort) FollowAheadXYR(o Object) {
	vp.xPos = o.xPos + (o.xSpd * (vp.width / 16))
	vp.yPos = o.yPos + (o.ySpd * (vp.height / 16))
	vp.rPos = o.rPos + (o.rSpd * (vp.height / 16))
}
