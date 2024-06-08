package game

import (
	"GoSnake/food"
	"GoSnake/sound"
	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	snake        *Snake
	food         *food.Food
	renderer     *Renderer
	logic        *GameLogic
	startManager *GameStartManager
	pauseManager *GamePauseManager
	audioManager *sound.AudioManager
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Updatable interface {
	Update() error
}

func NewGame(snake *Snake, food *food.Food, renderer *Renderer, logic *GameLogic, startManager *GameStartManager, pauseManager *GamePauseManager) *Game {
	return &Game{
		snake:        snake,
		food:         food,
		renderer:     renderer,
		logic:        logic,
		startManager: startManager,
		pauseManager: pauseManager,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.screen = screen
	g.renderer.drawBackground()
	g.renderer.drawSnake(g.snake.Body)
	g.renderer.drawFood(g.food.Position)
	gm := NewGameManager(g, g.startManager, g.pauseManager)
	g.renderer.drawUI(g.logic.score, g.logic.gameOver, g.logic.gameWon, g.startManager.IsGameStarted(), gm.gamePaused)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return vars.ScreenWidth, vars.ScreenHeight
}

func (g *Game) restart() {
	g.snake = NewSnake()
	g.food = food.NewFood()
	g.logic = NewGameLogic(g.audioManager)
}
