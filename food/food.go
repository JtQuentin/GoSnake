package food

import (
	"GoSnake/vars"
	"math/rand"
)

type Food struct {
	Position vars.Point
}

func NewFood() *Food {
	return &Food{
		Position: vars.Point{X: rand.Intn(vars.ScreenWidth / vars.TileSize), Y: rand.Intn(vars.ScreenHeight / vars.TileSize)},
	}
}

func (f *Food) Reset() {
	f.Position = vars.Point{X: rand.Intn(vars.ScreenWidth / vars.TileSize), Y: rand.Intn(vars.ScreenHeight / vars.TileSize)}
}
