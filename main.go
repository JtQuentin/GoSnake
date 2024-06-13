package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"

	"GoSnake/food"
	"GoSnake/game"
	"GoSnake/sound"
	"GoSnake/vars"
)

// main is the entry point of the application
func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a new audio context
	audioCtx, err := audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new audio manager
	audioManager := sound.NewAudioManager(audioCtx)

	// Initialize game components
	snake := game.NewSnake()
	food := food.NewFood()
	renderer := game.NewRenderer()
	logic := game.NewGameLogic(audioManager)
	gameStartManager := game.NewGameStartManager()
	gamePauseManager := game.NewGamePauseManager()

	// Create a new game instance
	g := game.NewGame(snake, food, renderer, logic, gameStartManager, gamePauseManager, audioManager)

	// Create a new game manager
	gameManager := game.NewGameManager(g, gameStartManager, gamePauseManager)

	// Set window size and title
	ebiten.SetWindowSize(vars.ScreenWidth*2, vars.ScreenHeight*2)
	ebiten.SetWindowTitle("GoSnake")

	// Run the game
	if err := ebiten.RunGame(gameManager); err != nil {
		log.Fatal(err)
	}

	// Close the audio manager
	audioManager.Close()
}
