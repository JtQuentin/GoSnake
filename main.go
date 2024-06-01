package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	audioCtx, err := audio.NewContext(44100)
	if err != nil {
		log.Fatal(err)
	}
	audioManager := NewAudioManager(audioCtx)

	snake := NewSnake()
	food := NewFood()
	renderer := NewRenderer()
	logic := NewGameLogic(audioManager)
	gameStartManager := NewGameStartManager()
	gamePauseManager := NewGamePauseManager()
	game := NewGame(snake, food, renderer, logic, gameStartManager, gamePauseManager)
	gameManager := NewGameManager(game, gameStartManager, gamePauseManager)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoSnake")
	if err := ebiten.RunGame(gameManager); err != nil {
		log.Fatal(err)
	}

	audioManager.Close()
}
