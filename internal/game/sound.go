package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	sounds     []*Sound
	sampleRate = 44100
	rcsSound   Sound
)

// Sound the game will automatically loop if not stopped
type Sound struct {
	start *audio.Player
	loop  *audio.Player
	stop  *audio.Player
}

// InitSounds initialize sounds
func InitSounds() {
	rcsSound = Sound{}

	audioContext := audio.NewContext(sampleRate)

	file, err := ebitenutil.OpenFile("../../assets/rcsstart.wav")
	decodedAudio, err := wav.Decode(audioContext, file)
	rcsSound.start, err = audio.NewPlayer(audioContext, decodedAudio)

	file, err = ebitenutil.OpenFile("../../assets/rcs.wav")
	decodedAudio, err = wav.Decode(audioContext, file)
	rcsSound.loop, err = audio.NewPlayer(audioContext, decodedAudio)

	file, err = ebitenutil.OpenFile("../../assets/rcsstop.wav")
	decodedAudio, err = wav.Decode(audioContext, file)
	rcsSound.stop, err = audio.NewPlayer(audioContext, decodedAudio)

	if err != nil {
		fmt.Println("derp")
		log.Fatal(err)
	}
}

// UpdateSounds to loop if needed
func UpdateSounds() {
	for _, sound := range sounds {
		if sound.loop.IsPlaying() && int(sound.loop.Current().Seconds()) == 4 {
			sound.loop.Rewind()
		}
	}
}

func startRcsSound() {
	rcsSound.loop.Play()
	rcsSound.start.Rewind()
	rcsSound.start.Play()
	sounds = append(sounds, &rcsSound)
}

func stopRcsSound() {
	for i := 0; i < len(sounds); i++ {
		if sounds[i] == &rcsSound {
			sounds[i] = sounds[len(sounds)-1]
			sounds = sounds[:len(sounds)-1]
		}
	}
	if rcsSound.loop.IsPlaying() && !rcsSound.stop.IsPlaying() {
		rcsSound.loop.Pause()
		rcsSound.stop.Rewind()
		rcsSound.stop.Play()
	}
}

// IsPlaying returns true if a sound effect is playing
func (sound *Sound) IsPlaying() bool {
	return sound.loop.IsPlaying()
}
