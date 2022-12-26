package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	rcsSoundFilePath            = "../../assets/rcs.wav"
	rcsOffSoundFilePath         = "../../assets/rcsoff.wav"
	missileSoundFilePath        = "../../assets/missile.wav"
	missileOffSoundFilePath     = "../../assets/missileoff.wav"
	missileReleaseSoundFilePath = "../../assets/missilerelease.wav"
	missleEmptySoundFilePath    = "../../assets/missileempty.wav"
)

var (
	sampleRate = 22050

	// Sound loops which can be heard at once.
	playerQueue   []*audio.Player
	missilePlayer *audio.Player
	rcsPlayer     *audio.Player

	// Sound effect players.
	rcsOffPlayer         *audio.Player
	missileOffPlayer     *audio.Player
	missileReleasePlayer *audio.Player
	missleEmptyPlayer    *audio.Player
)

// InitSounds initializes looping sounds.
func InitSounds() {
	audioContext := audio.NewContext(sampleRate)

	// Load the rcs sound.
	f, err := ebitenutil.OpenFile(rcsSoundFilePath)
	if err != nil {
		log.Fatal(err)
	}
	d, err := wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	sound := audio.NewInfiniteLoopWithIntro(d, 1*4*int64(sampleRate), 5*4*int64(sampleRate))
	rcsPlayer, err = audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	// Load the rcs off sound.
	f, err = ebitenutil.OpenFile(rcsOffSoundFilePath)
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	rcsOffPlayer, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile thrusting sound.
	f, err = ebitenutil.OpenFile(missileSoundFilePath)
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	sound = audio.NewInfiniteLoopWithIntro(d, 2*4*int64(sampleRate), 4*4*int64(sampleRate))
	missilePlayer, err = audio.NewPlayer(audioContext, sound)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile off sound.
	f, err = ebitenutil.OpenFile(missileOffSoundFilePath)
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	missileOffPlayer, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile release sound.
	f, err = ebitenutil.OpenFile(missileReleaseSoundFilePath)
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}
	missileReleasePlayer, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	// Load the missile warning sound.
	f, err = ebitenutil.OpenFile(missleEmptySoundFilePath)
	if err != nil {
		log.Fatal(err)
	}
	d, err = wav.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	missleEmptyPlayer, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}
}

// Queue the audio player.
func queuePlayer(p *audio.Player) {
	playerQueue = append(playerQueue, p)
	p.Rewind()
	p.Play()
}

// UnQueue the audio player.
// Keep looping if same audio player is in the audio queue.
func unQueuePlayer(p *audio.Player) {
	found := false
	for i := 0; i < len(playerQueue); {
		if playerQueue[i] != p {
			i++
			continue
		}
		if found {
			// Do not pause the sound, the same audio player is still in the queue.
			return
		}
		found = true
		// Remove the audio player from queue and replace it with the last one.
		playerQueue[i] = playerQueue[len(playerQueue)-1]
		playerQueue = playerQueue[:len(playerQueue)-1]
	}
	p.Pause()
}
