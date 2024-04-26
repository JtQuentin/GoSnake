package main

import (
	"math/rand"
)

type Food struct {
	Position Point
}

func NewFood() *Food {
	return &Food{
		Position: Point{X: rand.Intn(screenWidth / tileSize), Y: rand.Intn(screenHeight / tileSize)},
	}
}
