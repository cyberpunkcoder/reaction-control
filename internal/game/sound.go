package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	sampleRate = 22050
	looping    []*audio.Player
	missile    *audio.Player
	missileOff *audio.Player
	rcs        *audio.Player
	rcsOff     *audio.Player
	release    *audio.Player
	rocket     *audio.Player
)

// InitSounds initialize looping
func InitSounds() {
	audioContext := audio.NewContext(sampleRate)

	f, err := ebitenutil.OpenFile("../../assets/rcs.wav")
	d, err := wav.Decode(audioContext, f)
	sound := audio.NewInfiniteLoopWithIntro(d, 1*4*int64(sampleRate), 5*4*int64(sampleRate))
	rcs, err = audio.NewPlayer(audioContext, sound)

	f, err = ebitenutil.OpenFile("../../assets/missile.wav")
	d, err = wav.Decode(audioContext, f)
	sound = audio.NewInfiniteLoopWithIntro(d, 2*4*int64(sampleRate), 4*4*int64(sampleRate))
	missile, err = audio.NewPlayer(audioContext, sound)

	f, err = ebitenutil.OpenFile("../../assets/rcsoff.wav")
	d, err = wav.Decode(audioContext, f)
	rcsOff, err = audio.NewPlayer(audioContext, d)

	f, err = ebitenutil.OpenFile("../../assets/missileoff.wav")
	d, err = wav.Decode(audioContext, f)
	missileOff, err = audio.NewPlayer(audioContext, d)

	f, err = ebitenutil.OpenFile("../../assets/release.wav")
	d, err = wav.Decode(audioContext, f)
	release, err = audio.NewPlayer(audioContext, d)

	if err != nil {
		// There was a problem loading missile looping
		log.Fatal(err)
	}
}
