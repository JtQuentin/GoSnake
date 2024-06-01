package sound

import (
	"log"

	"os"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

func (am *AudioManager) PlayEatSound() {
	am.eatSoundPlayer.Rewind()
	am.eatSoundPlayer.Play()
}

func (am *AudioManager) PlayLoseSound() {
	am.loseSoundPlayer.Rewind()
	am.loseSoundPlayer.Play()
}

func loadAudioPlayer(ctx *audio.Context, filePath string) (*audio.Player, *os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	d, err := mp3.Decode(ctx, f)
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	p, err := audio.NewPlayer(ctx, d)
	if err != nil {
		f.Close()
		return nil, nil, err
	}

	return p, f, nil
}

type AudioManager struct {
	ctx             *audio.Context
	eatSoundPlayer  *audio.Player
	eatSoundFile    *os.File
	loseSoundPlayer *audio.Player
	loseSoundFile   *os.File
}

func NewAudioManager(ctx *audio.Context) *AudioManager {
	am := &AudioManager{ctx: ctx}
	var err error
	am.eatSoundPlayer, am.eatSoundFile, err = loadAudioPlayer(ctx, "eatSound.mp3")
	if err != nil {
		log.Fatal(err)
	}
	am.loseSoundPlayer, am.loseSoundFile, err = loadAudioPlayer(ctx, "loseSound.mp3")
	if err != nil {
		log.Fatal(err)
	}
	return am
}

func (am *AudioManager) Close() {
	am.eatSoundFile.Close()
	am.loseSoundFile.Close()
}
