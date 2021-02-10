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

	shipImage, _, err = ebitenutil.NewImageFromFile("../../assets/player.png")
	alienImage, _, err = ebitenutil.NewImageFromFile("../../assets/alien.png")
	missileImage, _, err = ebitenutil.NewImageFromFile("../../assets/missile.png")
	fusionImage, _, err = ebitenutil.NewImageFromFile("../../assets/fusion.png")
	rcsl, _, err = ebitenutil.NewImageFromFile("../../assets/rcsl.png")
	rcsr, _, err = ebitenutil.NewImageFromFile("../../assets/rcsr.png")
	rcsfl, _, err = ebitenutil.NewImageFromFile("../../assets/rcsfl.png")
	rcsfr, _, err = ebitenutil.NewImageFromFile("../../assets/rcsfr.png")
	rcsbl, _, err = ebitenutil.NewImageFromFile("../../assets/rcsbl.png")
	rcsbr, _, err = ebitenutil.NewImageFromFile("../../assets/rcsbr.png")
	space, _, err = ebitenutil.NewImageFromFile("../../assets/space.png")

	if err != nil {
		log.Fatal(err)
	}
}
