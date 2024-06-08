package food

import (
	"GoSnake/vars"
	"math/rand"
)

type Point struct {
	X, Y int
}

type Food struct {
	Position Point
}

func NewFood() *Food {
	return &Food{
		Position: Point{X: rand.Intn(vars.ScreenWidth / vars.TileSize), Y: rand.Intn(vars.ScreenHeight / vars.TileSize)},
	}
}

func (f *Food) Reset() {
	f.Position = Point{X: rand.Intn(vars.ScreenWidth / vars.TileSize), Y: rand.Intn(vars.ScreenHeight / vars.TileSize)}
}
