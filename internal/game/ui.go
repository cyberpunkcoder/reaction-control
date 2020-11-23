package game

import "github.com/hajimehoshi/ebiten/v2"

// ViewPort a player sees through
type ViewPort struct {
	Location
	width, height float64
}

// NewViewPort is initialized and returned
func NewViewPort(x float64, y float64, scale int) *ViewPort {
	w, h := ebiten.ScreenSizeInFullscreen()
	vp := ViewPort{width: float64(w / scale), height: float64(h / scale)}
	vp.x = x
	vp.y = y
	return &vp
}

// Follow object
func (vp *ViewPort) Follow(obj Object) {
	vp.x = obj.GetLocation().x
	vp.y = obj.GetLocation().y
}

// FollowAhead of object
func (vp *ViewPort) FollowAhead(obj Object) {
	vp.x = obj.GetLocation().x + (obj.GetPhysics().xSpd * (vp.width / 8))
	vp.y = obj.GetLocation().y + (obj.GetPhysics().ySpd * (vp.height / 8))
}
