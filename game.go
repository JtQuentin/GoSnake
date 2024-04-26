package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	snake        *Snake
	food         *Food
	renderer     *Renderer
	logic        *GameLogic
	startManager *GameStartManager
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Updatable interface {
	Update() error
}

type ScoreEntry struct {
	Name  string
	Score int
}

func NewGame(snake *Snake, food *Food, renderer *Renderer, logic *GameLogic, startManager *GameStartManager) *Game {
	return &Game{
		snake:        snake,
		food:         food,
		renderer:     renderer,
		logic:        logic,
		startManager: startManager,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.screen = screen
	g.renderer.drawBackground()
	g.renderer.drawSnake(g.snake.Body)
	g.renderer.drawFood(g.food.Position)
	g.renderer.drawUI(g.logic.score, g.logic.gameOver, g.logic.gameWon, g.startManager.IsGameStarted())
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) restart() {
	g.snake = NewSnake()
	g.food = NewFood()
	g.logic = NewGameLogic()
}
