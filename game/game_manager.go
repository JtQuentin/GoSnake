package game

import (
	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GameManager struct {
	game         *Game
	startManager *GameStartManager
	pauseManager *GamePauseManager
	gamePaused   bool
}

func NewGameManager(game *Game, startManager *GameStartManager, pauseManager *GamePauseManager) *GameManager {
	return &GameManager{game: game, startManager: startManager, pauseManager: pauseManager}
}

func (gm *GameManager) Update(screen *ebiten.Image) error {
	if !gm.startManager.IsGameStarted() {
		gm.startManager.HandleStartInput()
		return nil
	}

	if !gm.startManager.IsGameStarted() {
		gm.startManager.HandleStartInput()
		return nil
	}

	gm.pauseManager.HandlePauseInput()
	gm.gamePaused = gm.pauseManager.IsGamePaused()
	if gm.gamePaused {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		gm.game.restart()
	}

	if gm.game.logic.HandleGameState(inpututil.IsKeyJustPressed(ebiten.KeyR), gm.startManager.IsGameStarted()) {
		return nil
	}

	if gm.game.logic.UpdateTick() {
		gm.game.snake.updateDirection()
		gm.game.logic.CheckCollisions(gm.game.snake, gm.game.food)
	}

	gm.game.Draw(screen)

	return nil
}

func (gm *GameManager) Draw(screen *ebiten.Image) {
	gm.game.Draw(screen)
	gm.game.renderer.drawUI(gm.game.logic.score, gm.game.logic.gameOver, gm.game.logic.gameWon, gm.startManager.IsGameStarted(), gm.gamePaused)
}

func (gm *GameManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return vars.ScreenWidth, vars.ScreenHeight
}
