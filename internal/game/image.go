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

// InitImages initialize game images
func InitImages() {

	shipImage = mustLoadImageFromFile("../../assets/ship.png")
	alienImage = mustLoadImageFromFile("../../assets/alien.png")
	missileImage = mustLoadImageFromFile("../../assets/missile.png")
	fusionImage = mustLoadImageFromFile("../../assets/fusion.png")
	rcsl = mustLoadImageFromFile("../../assets/rcsl.png")
	rcsfr = mustLoadImageFromFile("../../assets/rcsfr.png")
	rcsfl = mustLoadImageFromFile("../../assets/rcsfl.png")
	rcsr = mustLoadImageFromFile("../../assets/rcsr.png")
	rcsbl = mustLoadImageFromFile("../../assets/rcsbl.png")
	rcsbr = mustLoadImageFromFile("../../assets/rcsbr.png")
	space = mustLoadImageFromFile("../../assets/space.png")
}

func mustLoadImageFromFile(imgPath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
