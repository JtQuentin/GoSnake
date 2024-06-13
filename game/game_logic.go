package game

import (
	"GoSnake/food"
	"GoSnake/sound"
	"GoSnake/vars"
)

// GameLogic represents the game's logic
type GameLogic struct {
	score         int                 // The player's current score
	gameOver      bool                // Whether the game is over
	gameWon       bool                // Whether the player has won the game
	speed         int                 // The game's speed, which affects the update rate
	updateCounter int                 // A counter used to control the update rate
	audioManager  *sound.AudioManager // A pointer to an AudioManager object, which handles sound effects
}

// NewGameLogic creates a new GameLogic object with default values
func NewGameLogic(audioManager *sound.AudioManager) *GameLogic {
	return &GameLogic{
		speed:        10,           // Initial game speed
		audioManager: audioManager, // AudioManager for playing sounds
	}
}

// HandleGameState checks the game state and handles restarts
func (gl *GameLogic) HandleGameState(restartPressed, gameStarted bool) bool {
	// If the game hasn't started or it's over or won, and restart is pressed, restart the game
	if !gameStarted || gl.gameOver || gl.gameWon {
		if restartPressed {
			gl.restartGame()
			return false
		}
		return true
	}
	return false
}

// restartGame resets the game state, including the score, game over flag, and speed
func (gl *GameLogic) restartGame() {
	gl.score = 0
	gl.gameOver = false
	gl.gameWon = false
	gl.speed = 10
	gl.updateCounter = 0
}

// UpdateTick increments the update counter and checks if it's time to update the game state
func (gl *GameLogic) UpdateTick() bool {
	gl.updateCounter++
	// If the update counter is less than the speed, don't update the game state
	if gl.updateCounter < gl.speed {
		return false
	}
	gl.updateCounter = 0
	return true
}

// CheckCollisions checks for collisions between the snake and the food or the game boundaries
func (gl *GameLogic) CheckCollisions(snake *Snake, food *food.Food) {
	head := snake.Body[0]
	// Check for collision with game boundaries
	if head.X < 0 || head.Y < 0 || head.X >= vars.ScreenWidth/vars.TileSize || head.Y >= vars.ScreenHeight/vars.TileSize {
		gl.gameOver = true
		gl.speed = 10
		SaveScore(gl.score)
		if gl.audioManager != nil {
			gl.audioManager.PlayLoseSound()
		}
		return
	}

	// Check for self-collisions
	for _, part := range snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			gl.gameOver = true
			gl.speed = 10
			if gl.audioManager != nil {
				gl.audioManager.PlayLoseSound()
			}
			return
		}
	}

	// Check for collision with food
	if head.X == food.Position.X && head.Y == food.Position.Y {
		gl.score++
		snake.GrowCounter += 1
		food.Reset()
		if gl.audioManager != nil {
			gl.audioManager.PlayEatSound()
		}

		// Check if the player has won the game
		if gl.score == 25 {
			gl.gameWon = true
			gl.speed = 10
		} else {
			// Decrease the game speed if it's greater than 2
			if gl.speed > 2 {
				gl.speed--
			}
		}
	}
}
