package game

import (
	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Snake struct represents the snake in the game
type Snake struct {
	Body        []vars.Point // Body is a slice of points that represents the body of the snake
	Direction   vars.Point   // Direction is the current direction of the snake
	GrowCounter int          // GrowCounter is the number of times the snake needs to grow
}

// NewSnake function creates a new snake and returns a pointer to it
func NewSnake() *Snake {
	return &Snake{
		Body:      []vars.Point{{X: vars.ScreenWidth / vars.TileSize / 2, Y: vars.ScreenHeight / vars.TileSize / 2}}, // Initialize the snake in the middle of the screen
		Direction: vars.Point{X: 1, Y: 0},                                                                            // The snake starts moving to the right
	}
}

// processInput processes the key inputs and sets the new direction
func (s *Snake) processInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if s.Direction.X == 0 {
			s.Direction = vars.Point{X: -1, Y: 0} // Move left
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if s.Direction.X == 0 {
			s.Direction = vars.Point{X: 1, Y: 0} // Move right
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if s.Direction.Y == 0 {
			s.Direction = vars.Point{X: 0, Y: -1} // Move up
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if s.Direction.Y == 0 {
			s.Direction = vars.Point{X: 0, Y: 1} // Move down
		}
	}
}

// updateDirection updates the direction of the snake based on the processed input
func (s *Snake) updateDirection() {
	s.Move() // Move the snake in the updated direction
}

// Move function moves the snake in the current direction
func (s *Snake) Move() {
	newHead := vars.Point{X: s.Body[0].X + s.Direction.X, Y: s.Body[0].Y + s.Direction.Y} // Calculate the new head of the snake
	s.Body = append([]vars.Point{newHead}, s.Body...)                                     // Add the new head to the body of the snake
	if s.GrowCounter > 0 {
		s.GrowCounter-- // If the snake needs to grow, decrease the grow counter
	} else {
		s.Body = s.Body[:len(s.Body)-1] // If the snake doesn't need to grow, remove the last element of the body
	}
}
