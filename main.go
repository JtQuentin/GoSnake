package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
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

// Function to update the Snake's position
func (g *Game) Update(screen *ebiten.Image) error {

	g.updateCounter++
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0

	g.snake.Move()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.Direction.X == 0 {
		g.snake.Direction = Point{X: -1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.Direction.X == 0 {
		g.snake.Direction = Point{X: 1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.Direction.Y == 0 {
		g.snake.Direction = Point{X: 0, Y: -1}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.Direction.Y == 0 {
		g.snake.Direction = Point{X: 0, Y: 1}
	}

	head := g.snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		g.gameOver = true
		g.speed = 10
	}

	for _, part := range g.snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			g.gameOver = true
			g.speed = 10
		}
	}

	if head.X == g.food.Position.X && head.Y == g.food.Position.Y {
		g.score++
		g.snake.GrowCounter += 1
		g.food = NewFood()

		if g.speed > 2 {
			g.speed--
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(color.RGBA{154, 198, 0, 255})

	// Draw snake
	for _, p := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(p.X*tileSize), float64(p.Y*tileSize), tileSize, tileSize, color.RGBA{33, 50, 15, 255})
	}

	// Draw food
	ebitenutil.DrawRect(screen, float64(g.food.Position.X*tileSize), float64(g.food.Position.Y*tileSize), tileSize, tileSize, color.RGBA{231, 71, 29, 255})

	// Create a font.Face
	face := basicfont.Face7x13

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, face, 5, screenHeight-5, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		snake:    NewSnake(),
		food:     NewFood(),
		gameOver: false,
		gameWon:  false,
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
