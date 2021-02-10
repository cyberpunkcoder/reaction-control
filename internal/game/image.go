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

	mustLoadImageFromFile(shipImage, "../../assets/player.png")
	mustLoadImageFromFile(alienImage, "../../assets/alien.png")
	mustLoadImageFromFile(missileImage, "../../assets/missile.png")
	mustLoadImageFromFile(fusionImage, "../../assets/fusion.png")
	mustLoadImageFromFile(rcsl, "../../assets/rcsl.png")
	mustLoadImageFromFile(rcsr, "../../assets/rcsr.png")
	mustLoadImageFromFile(rcsbl, "../../assets/rcsbl.png")
	mustLoadImageFromFile(rcsbr, "../../assets/rcsbr.png")
	mustLoadImageFromFile(space, "../../assets/space.png")
}

func mustLoadImageFromFile(img *ebiten.Image, imgPath string) {
	img, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}
}
