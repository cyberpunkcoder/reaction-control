package game

import "github.com/hajimehoshi/ebiten/v2"

// ViewPort a player sees through
type ViewPort struct {
	xPos, yPos, width, height float64
}

// NewViewPort is initialized and returned
func NewViewPort(xpos float64, ypos float64, scale int) *ViewPort {
	w, h := ebiten.ScreenSizeInFullscreen()
	return &ViewPort{float64(w / 2), float64(h / 2), float64(w / scale), float64(h / scale)}
}
