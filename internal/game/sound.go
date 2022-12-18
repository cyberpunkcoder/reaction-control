package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	sampleRate = 22050

	// Sound loops which can be heard at once.
	queue   []*audio.Player
	missile *audio.Player
	rcs     *audio.Player

	// Sound effect players.
	missileOff *audio.Player
	rcsOff     *audio.Player
	release    *audio.Player
	warning    *audio.Player
)

// InitSounds initializes looping sounds.
func InitSounds() {
	audioContext := audio.NewContext(sampleRate)

	// Load the rcs sound.
	f, err := ebitenutil.OpenFile("../../assets/rcs.wav")
	if err != nil {
		log.Fatal(err)
	}
	d, err := wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	sound := audio.NewInfiniteLoopWithIntro(d, 1*4*int64(sampleRate), 5*4*int64(sampleRate))
	rcs, err = audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	// Load the rcs off sound.
	f, err = ebitenutil.OpenFile("../../assets/rcsoff.wav")
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	rcsOff, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile thrusting sound.
	f, err = ebitenutil.OpenFile("../../assets/missile.wav")
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	sound = audio.NewInfiniteLoopWithIntro(d, 2*4*int64(sampleRate), 4*4*int64(sampleRate))
	missile, err = audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile off sound.
	f, err = ebitenutil.OpenFile("../../assets/missileoff.wav")
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	missileOff, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile release sound.
	f, err = ebitenutil.OpenFile("../../assets/release.wav")
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	release, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile warning sound.
	f, err = ebitenutil.OpenFile("../../assets/warning.wav")
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	warning, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}
}

// Queue the audio player.
func queuePlayer(p *audio.Player) {
	queue = append(queue, p)
	p.Rewind()
	p.Play()
}

// UnQueue the audio player. 
// Keep looping if same audio player is in the audio queue.
func unQueuePlayer(p *audio.Player) {
	found := false
	for i := 0; i < len(queue); {
		if queue[i] != p {
			i++
			continue
		}
		if found {
			// Do not pause the sound, the same audio player is still in the queue.
			return
		}
		found = true
		// Remove the audio player from queue and replace it with the last one.
		queue[i] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
	}
	p.Pause()
}
