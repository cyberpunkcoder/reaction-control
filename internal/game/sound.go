package game

import (
	"bytes"
	"log"

	rice "github.com/GeertJohan/go.rice"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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
func InitSounds(soundBox *rice.Box) {
	audioContext := audio.NewContext(sampleRate)

	d := mustLoadSoundAsStream(soundBox, "rcs.wav", audioContext)
	sound := audio.NewInfiniteLoopWithIntro(d, 1*4*int64(sampleRate), 5*4*int64(sampleRate))
	rcs, err = audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	d = mustLoadSoundAsStream(soundBox, "missile.wav", audioContext)
	sound = audio.NewInfiniteLoopWithIntro(d, 2*4*int64(sampleRate), 4*4*int64(sampleRate))
	missile, err = audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	d = mustLoadSoundAsStream(soundBox, "rcsoff.wav", audioContext)
	rcsOff, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	d = mustLoadSoundAsStream(soundBox, "missileoff.wav", audioContext)
	missileOff, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	d = mustLoadSoundAsStream(soundBox, "release.wav", audioContext)
	release, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	d = mustLoadSoundAsStream(soundBox, "warning.wav", audioContext)
	warning, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}
}

func mustLoadSoundAsStream(soundBox *rice.Box, soundFileName string, audioContext *audio.Context) *wav.Stream {
	soundFile := bytes.NewReader(soundBox.MustBytes(soundFileName))
	d, err := wav.Decode(audioContext, soundFile)
	if err != nil {
		log.Fatalf("Unable to decode sound %s: %v\n", soundFileName, err)
	}

	return d
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
