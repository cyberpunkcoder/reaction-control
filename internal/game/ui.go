package game

import "github.com/hajimehoshi/ebiten/v2"

// ViewPort a player sees through
type ViewPort struct {
	Location
	width, height float64
}

// NewViewPort is initialized and returned
func NewViewPort(x float64, y float64) *ViewPort {
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
	vp.x = obj.GetLocation().x + (obj.GetSpeed().xSpd * (vp.width / 16))
	vp.y = obj.GetLocation().y + (obj.GetSpeed().ySpd * (vp.height / 16))
}
