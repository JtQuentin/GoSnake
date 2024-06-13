package game

import (
	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// GameManager manages the game state and user input
type GameManager struct {
	game         *Game
	startManager *GameStartManager // Manages the game start state
	pauseManager *GamePauseManager // Manages the game pause state
	gamePaused   bool              // Indicates if the game is paused
}

// NewGameManager creates a new GameManager object
func NewGameManager(game *Game, startManager *GameStartManager, pauseManager *GamePauseManager) *GameManager {
	return &GameManager{game: game, startManager: startManager, pauseManager: pauseManager}
}

// Update updates the game state and handles user input
func (gm *GameManager) Update(screen *ebiten.Image) error {
	// If the game has not started, handle start input
	if !gm.startManager.IsGameStarted() {
		gm.startManager.HandleStartInput()
		return nil
	}

	// If the 'R' key is pressed, restart the game
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		gm.game.restart()
	}

	// Handle pause input only if the game is not over
	if !gm.game.logic.gameOver {
		gm.pauseManager.HandlePauseInput()
		gm.gamePaused = gm.pauseManager.IsGamePaused()
	}

	// If the game is paused or over, return
	if gm.gamePaused || gm.game.logic.gameOver {
		return nil
	}

	// Handle game state and restart if necessary
	if gm.game.logic.HandleGameState(inpututil.IsKeyJustPressed(ebiten.KeyR), gm.startManager.IsGameStarted()) {
		return nil
	}

	// Process input for the snake direction
	gm.game.snake.processInput()

	// Update the game logic and check for collisions
	if gm.game.logic.UpdateTick() {
		gm.game.snake.updateDirection()
		gm.game.logic.CheckCollisions(gm.game.snake, gm.game.food)
	}

	// Draw the game
	gm.game.Draw(screen)

	return nil
}

// Draw draws the game and UI
func (gm *GameManager) Draw(screen *ebiten.Image) {
	// Draw the game
	gm.game.Draw(screen)
	// Draw the UI
	gm.game.renderer.drawUI(gm.game.logic.score, gm.game.logic.gameOver, gm.game.logic.gameWon, gm.startManager.IsGameStarted(), gm.gamePaused)
}

// Layout returns the screen width and height
func (gm *GameManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the screen dimensions
	return vars.ScreenWidth, vars.ScreenHeight
}
