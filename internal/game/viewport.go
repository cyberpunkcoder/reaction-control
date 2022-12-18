package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// ViewPort is a window a player sees through.
type ViewPort struct {
	Position
	Speed
	zoom int
	width, height float64
}

// NewViewPort initializes a new view port.
func NewViewPort(p Position) *ViewPort {
	w, h := ebiten.ScreenSizeInFullscreen()

	// Scale up the game for monitor size.
	zoom := 2

	if w > 1920 {
		zoom = 4
	} else if w > 1024 {
		zoom = 3
	}

	vp := ViewPort{
		width:  float64(w / zoom),
		height: float64(h / zoom),
		zoom:   zoom,
		Position: Position{
			xPos: p.xPos,
			yPos: p.yPos,
		},
	}
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

// Max position visible on viewport
func (vp *ViewPort) Max() (float64, float64) {
	w, h := vp.width/2, vp.height/2
	radAng := vp.rPos * (math.Pi / 180)

	x := math.Abs(w*math.Cos(radAng)) - math.Min(h*math.Sin(radAng), -h*math.Sin(radAng))
	y := math.Abs(w*math.Sin(radAng)) + math.Max(h*math.Cos(radAng), -h*math.Cos(radAng))

	return x, y
}
