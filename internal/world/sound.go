package world

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	rcsSound      *audio.Player
	rcsStartSound *audio.Player
	rcsStopSound  *audio.Player
	sampleRate    = 44100
)

func InitSounds() {
	audioContext := audio.NewContext(sampleRate)

	file, err := ebitenutil.OpenFile("../../assets/rcs.wav")
	decodedAudio, err := wav.Decode(audioContext, file)
	rcsSound, err = audio.NewPlayer(audioContext, decodedAudio)

	file, err = ebitenutil.OpenFile("../../assets/rcsstart.wav")
	decodedAudio, err = wav.Decode(audioContext, file)
	rcsStartSound, err = audio.NewPlayer(audioContext, decodedAudio)

	file, err = ebitenutil.OpenFile("../../assets/rcsstop.wav")
	decodedAudio, err = wav.Decode(audioContext, file)
	rcsStopSound, err = audio.NewPlayer(audioContext, decodedAudio)

	if err != nil {
		fmt.Println("derp")
		log.Fatal(err)
	}
}

func UpdateSound() {
	if rcsSound.IsPlaying() && int(rcsSound.Current().Seconds()) == 4 {
		rcsSound.Rewind()
	}
}

func startRcsSound() {
	rcsSound.Play()
	rcsStartSound.Rewind()
	rcsStartSound.Play()
}

func stopRcsSound() {
	if rcsSound.IsPlaying() && !rcsStopSound.IsPlaying() {
		rcsSound.Pause()
		rcsStopSound.Rewind()
		rcsStopSound.Play()
	}
}
