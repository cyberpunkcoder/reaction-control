package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	sampleRate   = 44100
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

//loop an audio player and pause other loops of the same track to prevent distortion
func loop(p *audio.Player) {
	for _, loop := range looping {
		if loop == p {
			loop.Pause()
			break
		}
	}
	looping = append(looping, p)
	p.Rewind()
	p.Play()
}

// Stop a player that is being looped and resume other loops of the same track
func stoploop(p *audio.Player) {
	found := false
	for i := 0; i < len(looping); i++ {
		if looping[i] == p {
			if p.IsPlaying() && !found {
				looping[i].Pause()
				looping[i] = looping[len(looping)-1]
				looping = looping[:len(looping)-1]
				found = true
			} else {
				looping[i].Play()
				break
			}
		}
	}
}
