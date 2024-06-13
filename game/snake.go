package game

import (
	"GoSnake/food"
	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
)

// Snake struct represents the snake in the game
type Snake struct {
	Body        []food.Point // Body is a slice of points that represents the body of the snake
	Direction   food.Point   // Direction is the current direction of the snake
	GrowCounter int          // GrowCounter is the number of times the snake needs to grow
}

// NewSnake function creates a new snake and returns a pointer to it
func NewSnake() *Snake {
	return &Snake{
		Body:      []food.Point{{X: vars.ScreenWidth / vars.TileSize / 2, Y: vars.ScreenHeight / vars.TileSize / 2}}, // Initialize the snake in the middle of the screen
		Direction: food.Point{X: 1, Y: 0},                                                                            // The snake starts moving to the right
	}
}

// updateDirection function updates the direction of the snake based on the key pressed
func (s *Snake) updateDirection() {
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)) && s.Direction.X == 0 {
		s.Direction = food.Point{X: -1, Y: 0} // Move left
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)) && s.Direction.X == 0 {
		s.Direction = food.Point{X: 1, Y: 0} // Move right
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)) && s.Direction.Y == 0 {
		s.Direction = food.Point{X: 0, Y: -1} // Move up
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)) && s.Direction.Y == 0 {
		s.Direction = food.Point{X: 0, Y: 1} // Move down
	}
	s.Move() // Move the snake in the updated direction
}

// Move function moves the snake in the current direction
func (s *Snake) Move() {
	newHead := food.Point{X: s.Body[0].X + s.Direction.X, Y: s.Body[0].Y + s.Direction.Y} // Calculate the new head of the snake
	s.Body = append([]food.Point{newHead}, s.Body...)                                     // Add the new head to the body of the snake
	if s.GrowCounter > 0 {
		s.GrowCounter-- // If the snake needs to grow, decrease the grow counter
	} else {
		s.Body = s.Body[:len(s.Body)-1] // If the snake doesn't need to grow, remove the last element of the body
	}
}
