package game

import (
	"GoSnake/food"

	"github.com/hajimehoshi/ebiten"
)

type Snake struct {
	Body        []food.Point
	Direction   food.Point
	GrowCounter int
}

func NewSnake() *Snake {
	return &Snake{
		Body:      []food.Point{{X: ScreenWidth / TileSize / 2, Y: ScreenHeight / TileSize / 2}},
		Direction: food.Point{X: 1, Y: 0},
	}
}

func (s *Snake) updateDirection() {
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)) && s.Direction.X == 0 {
		s.Direction = food.Point{X: -1, Y: 0}
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)) && s.Direction.X == 0 {
		s.Direction = food.Point{X: 1, Y: 0}
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)) && s.Direction.Y == 0 {
		s.Direction = food.Point{X: 0, Y: -1}
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)) && s.Direction.Y == 0 {
		s.Direction = food.Point{X: 0, Y: 1}
	}
	s.Move()
}

func (s *Snake) Move() {
	newHead := food.Point{X: s.Body[0].X + s.Direction.X, Y: s.Body[0].Y + s.Direction.Y}
	s.Body = append([]food.Point{newHead}, s.Body...)
	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}
