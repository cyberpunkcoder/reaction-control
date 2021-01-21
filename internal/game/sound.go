package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	sampleRate   = 22050
	looping      []*audio.Player
	missileSound *audio.Player
	rcs          *audio.Player
)

// InitSounds initialize looping
func InitSounds() {
	audioContext := audio.NewContext(sampleRate)

	f, err := os.Open("../../assets/rcs.wav")
	rcsAudio := audio.NewInfiniteLoopWithIntro(f, 1*4*int64(sampleRate), 5*4*int64(sampleRate))

	rcs, err = audio.NewPlayer(audioContext, rcsAudio)

	if err != nil {
		// There was a problem loading missile looping
		log.Fatal(err)
	}
}
