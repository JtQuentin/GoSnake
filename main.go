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
	audioManager := sound.NewAudioManager(audioCtx)

	snake := game.NewSnake()
	food := food.NewFood()
	renderer := game.NewRenderer()
	logic := game.NewGameLogic(audioManager)
	gameStartManager := game.NewGameStartManager()
	gamePauseManager := game.NewGamePauseManager()
	g := game.NewGame(snake, food, renderer, logic, gameStartManager, gamePauseManager)
	gameManager := game.NewGameManager(g, gameStartManager, gamePauseManager)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoSnake")
	if err := ebiten.RunGame(gameManager); err != nil {
		log.Fatal(err)
	}

	audioManager.Close()
}
