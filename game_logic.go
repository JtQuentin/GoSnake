package main

type GameLogic struct {
	score         int
	gameOver      bool
	gameWon       bool
	speed         int
	updateCounter int
}

func NewGameLogic() *GameLogic {
	return &GameLogic{speed: 10}
}

func (gl *GameLogic) HandleGameState(restartPressed, gameStarted bool) bool {
	if !gameStarted || gl.gameOver || gl.gameWon {
		if restartPressed {
			gl.restartGame()
			return false
		}
		return true
	}
	return false
}

func (gl *GameLogic) restartGame() {
	gl.score = 0
	gl.gameOver = false
	gl.gameWon = false
	gl.speed = 10
	gl.updateCounter = 0
}

func (gl *GameLogic) UpdateTick() bool {
	gl.updateCounter++
	if gl.updateCounter < gl.speed {
		return false
	}
	gl.updateCounter = 0
	return true
}

func (gl *GameLogic) CheckCollisions(snake *Snake, food *Food) {
	head := snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		gl.gameOver = true
		gl.speed = 10
		SaveScore(gl.score)
		return
	}

	for _, part := range snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			gl.gameOver = true
			gl.speed = 10
			return
		}
	}

	if head.X == food.Position.X && head.Y == food.Position.Y {
		gl.score++
		snake.GrowCounter += 1
		*food = *NewFood()

		if gl.score == 25 {
			gl.gameWon = true
			gl.speed = 10
		} else {
			if gl.speed > 2 {
				gl.speed--
			}
		}
	}
}
