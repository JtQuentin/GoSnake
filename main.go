package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	snake := NewSnake()
	food := NewFood()
	renderer := NewRenderer()
	logic := NewGameLogic()
	gameStartManager := NewGameStartManager()
	gamePauseManager := NewGamePauseManager()
	game := NewGame(snake, food, renderer, logic, gameStartManager)
	gameManager := NewGameManager(game, gameStartManager, gamePauseManager)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoSnake")
	if err := ebiten.RunGame(gameManager); err != nil {
		log.Fatal(err)
	}
}
