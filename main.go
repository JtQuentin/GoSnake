package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

type Game struct {
	snake         *Snake
	food          *Food
	score         int
	gameOver      bool
	gameWon       bool
	ticks         int
	updateCounter int
	speed         int
}

type Point struct {
	X int
	Y int
}

type Snake struct {
	Body        []Point
	Direction   Point
	GrowCounter int
}

type Food struct {
	Position Point
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw snake
	for _, p := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(p.X*tileSize), float64(p.Y*tileSize), tileSize, tileSize, color.RGBA{0, 255, 0, 255})
	}

	// Draw food
	ebitenutil.DrawRect(screen, float64(g.food.Position.X*tileSize), float64(g.food.Position.Y*tileSize), tileSize, tileSize, color.RGBA{255, 0, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		snake:    NewSnake(),
		food:     NewFood(),
		gameOver: false,
		ticks:    0,
		speed:    10,
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoSnake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// Function to create a Snake
func NewSnake() *Snake {
	return &Snake{
		Body: []Point{
			{X: screenWidth / tileSize / 2, Y: screenHeight / tileSize / 2},
		},
		Direction: Point{X: 1, Y: 0},
	}
}

// Function to move the Snake
func (s *Snake) Move() {
	newHead := Point{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}
	s.Body = append([]Point{newHead}, s.Body...)

	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

// Function to create the Food
func NewFood() *Food {
	return &Food{
		Position: Point{
			X: rand.Intn(screenWidth / tileSize),
			Y: rand.Intn(screenHeight / tileSize),
		},
	}
}
