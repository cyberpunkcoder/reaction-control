package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	sampleRate = 22050

	// Sound loops
	queue   []*audio.Player
	missile *audio.Player
	rcs     *audio.Player
	rocket  *audio.Player

	// Sound effects
	missileOff *audio.Player
	rcsOff     *audio.Player
	release    *audio.Player
	warning    *audio.Player
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

	f, err = ebitenutil.OpenFile("../../assets/warning.wav")
	d, err = wav.Decode(audioContext, f)
	warning, err = audio.NewPlayer(audioContext, d)

	if err != nil {
		// There was a problem loading missile looping
		log.Fatal(err)
	}
}

// Queue audio player
func queuePlayer(p *audio.Player) {
	queue = append(queue, p)
	p.Rewind()
	p.Play()
}

// UnQueue audio player, keep looping if same player is in queue
func unQueuePlayer(p *audio.Player) {
	found := false
	for i := 0; i < len(queue); {
		if queue[i] != p {
			i++
			continue
		}
		if found {
			// Do not pause the sound, the same player is still in queue
			return
		}
		found = true
		// Remove player from queue and replace it with the last player
		queue[i] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
	}
	p.Pause()
}
