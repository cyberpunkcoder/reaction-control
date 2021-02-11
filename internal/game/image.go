package game

import (
	"bytes"
	"image"
	_ "image/png" // Required for ebitenutil.NewImageFromFile()
	"log"

	rice "github.com/GeertJohan/go.rice"
	"github.com/hajimehoshi/ebiten/v2"
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
func InitImages(imgBox *rice.Box) {
	shipImage = mustLoadImage(imgBox, "ship.png")
	alienImage = mustLoadImage(imgBox, "alien.png")
	missileImage = mustLoadImage(imgBox, "missile.png")
	fusionImage = mustLoadImage(imgBox, "alien.png")
	rcsl = mustLoadImage(imgBox, "rcsl.png")
	rcsfr = mustLoadImage(imgBox, "rcsfr.png")
	rcsfl = mustLoadImage(imgBox, "rcsfl.png")
	rcsr = mustLoadImage(imgBox, "rcsr.png")
	rcsbl = mustLoadImage(imgBox, "rcsbl.png")
	rcsbr = mustLoadImage(imgBox, "rcsbr.png")
	space = mustLoadImage(imgBox, "space.png")
}

func mustLoadImage(imgBox *rice.Box, imgFileName string) *ebiten.Image {
	imgBytes := bytes.NewReader(imgBox.MustBytes(imgFileName))

	img, _, err := image.Decode(imgBytes)
	if err != nil {
		log.Fatalf("Unable to decode image %s: %v\n", imgFileName, err)
	}

	img2 := ebiten.NewImageFromImage(img)

	return img2
}
