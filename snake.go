package main

import "github.com/hajimehoshi/ebiten"

type Point struct {
	X, Y int
}

type Snake struct {
	Body        []Point
	Direction   Point
	GrowCounter int
}

func NewSnake() *Snake {
	return &Snake{
		Body:      []Point{{X: screenWidth / tileSize / 2, Y: screenHeight / tileSize / 2}},
		Direction: Point{X: 1, Y: 0},
	}
}

func (s *Snake) updateDirection() {
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)) && s.Direction.X == 0 {
		s.Direction = Point{X: -1, Y: 0}
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)) && s.Direction.X == 0 {
		s.Direction = Point{X: 1, Y: 0}
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)) && s.Direction.Y == 0 {
		s.Direction = Point{X: 0, Y: -1}
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)) && s.Direction.Y == 0 {
		s.Direction = Point{X: 0, Y: 1}
	}
	s.Move()
}

func (s *Snake) Move() {
	newHead := Point{X: s.Body[0].X + s.Direction.X, Y: s.Body[0].Y + s.Direction.Y}
	s.Body = append([]Point{newHead}, s.Body...)
	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}
