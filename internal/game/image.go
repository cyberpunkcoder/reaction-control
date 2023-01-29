package game

import (
	_ "image/png" // Required for ebitenutil.NewImageFromFile()
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	shipImage    *ebiten.Image
	alienImage   *ebiten.Image
	fusionImage  *ebiten.Image
	missileImage *ebiten.Image
	rcsl         *ebiten.Image
	rcsr         *ebiten.Image
	rcsfl        *ebiten.Image
	rcsfr        *ebiten.Image
	rcsbl        *ebiten.Image
	rcsbr        *ebiten.Image
	space        *ebiten.Image
)

// InitImages initializes the game images.
func InitImages() {
	shipImage = mustLoadImageFromFile("../../assets/images/ship.png")
	alienImage = mustLoadImageFromFile("../../assets/images/alien.png")
	missileImage = mustLoadImageFromFile("../../assets/images/missile.png")
	fusionImage = mustLoadImageFromFile("../../assets/images/fusion.png")
	rcsl = mustLoadImageFromFile("../../assets/images/rcsl.png")
	rcsfr = mustLoadImageFromFile("../../assets/images/rcsfr.png")
	rcsfl = mustLoadImageFromFile("../../assets/images/rcsfl.png")
	rcsr = mustLoadImageFromFile("../../assets/images/rcsr.png")
	rcsbl = mustLoadImageFromFile("../../assets/images/rcsbl.png")
	rcsbr = mustLoadImageFromFile("../../assets/images/rcsbr.png")
	space = mustLoadImageFromFile("../../assets/images/space.png")
}

// mustLoadImageFromFile loads an image from a file and panics if it fails.
func mustLoadImageFromFile(imgPath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
