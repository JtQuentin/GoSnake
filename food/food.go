package food

import (
	"math/rand"
)

type Point struct {
	X, Y int
}

const (
	ScreenWidth  = 320
	ScreenHeight = 240
	TileSize     = 5
)

type Food struct {
	Position Point
}

func NewFood() *Food {
	return &Food{
		Position: Point{X: rand.Intn(ScreenWidth / TileSize), Y: rand.Intn(ScreenHeight / TileSize)},
	}
}

func (f *Food) Reset() {
	f.Position = Point{X: rand.Intn(ScreenWidth / TileSize), Y: rand.Intn(ScreenHeight / TileSize)}
}
