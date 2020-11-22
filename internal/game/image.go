package game

import (
	_ "image/png" // Required for ebitenutil.NewImageFromFile()
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	shipImage *ebiten.Image
	rcsfl     *ebiten.Image
	rcsfr     *ebiten.Image
	rcsbl     *ebiten.Image
	rcsbr     *ebiten.Image
	space     *ebiten.Image
)

// InitImages initialize game images
func InitImages() {

	shipImage, _, err = ebitenutil.NewImageFromFile("../../assets/player.png")
	rcsfl, _, err = ebitenutil.NewImageFromFile("../../assets/rcsfl.png")
	rcsfr, _, err = ebitenutil.NewImageFromFile("../../assets/rcsfr.png")
	rcsbl, _, err = ebitenutil.NewImageFromFile("../../assets/rcsbl.png")
	rcsbr, _, err = ebitenutil.NewImageFromFile("../../assets/rcsbr.png")
	space, _, err = ebitenutil.NewImageFromFile("../../assets/space.png")

	if err != nil {
		log.Fatal(err)
	}
}
