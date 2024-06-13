package sound

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

// AudioManager represents an audio manager
type AudioManager struct {
	ctx             *audio.Context // The audio context
	eatSoundPlayer  *audio.Player  // The audio player for the eat sound
	eatSoundFile    *os.File       // The file for the eat sound
	loseSoundPlayer *audio.Player  // The audio player for the lose sound
	loseSoundFile   *os.File       // The file for the lose sound
}

// NewAudioManager creates a new AudioManager object
func NewAudioManager(ctx *audio.Context) *AudioManager {
	am := &AudioManager{ctx: ctx}
	var err error
	// Load the eat sound
	am.eatSoundPlayer, am.eatSoundFile, err = loadAudioPlayer(ctx, "sound/eatSound.mp3")
	if err != nil {
		log.Fatal(err)
	}
	// Load the lose sound
	am.loseSoundPlayer, am.loseSoundFile, err = loadAudioPlayer(ctx, "sound/loseSound.mp3")
	if err != nil {
		log.Fatal(err)
	}
	return am
}

// PlayEatSound plays the eat sound
func (am *AudioManager) PlayEatSound() {
	am.eatSoundPlayer.Rewind() // Rewind the audio player to the start
	am.eatSoundPlayer.Play()   // Play the audio
}

// PlayLoseSound plays the lose sound
func (am *AudioManager) PlayLoseSound() {
	am.loseSoundPlayer.Rewind() // Rewind the audio player to the start
	am.loseSoundPlayer.Play()   // Play the audio
}

// loadAudioPlayer loads an audio player from a file
func loadAudioPlayer(ctx *audio.Context, filePath string) (*audio.Player, *os.File, error) {
	f, err := os.Open(filePath) // Open the audio file
	if err != nil {
		return nil, nil, err
	}

	d, err := mp3.Decode(ctx, f) // Decode the MP3 file
	if err != nil {
		f.Close() // Close the file if there is an error
		return nil, nil, err
	}

	p, err := audio.NewPlayer(ctx, d) // Create a new audio player
	if err != nil {
		f.Close() // Close the file if there is an error
		return nil, nil, err
	}

	return p, f, nil // Return the audio player and the file
}

// Close closes the audio files
func (am *AudioManager) Close() {
	am.eatSoundFile.Close()  // Close the eat sound file
	am.loseSoundFile.Close() // Close the lose sound file
}
